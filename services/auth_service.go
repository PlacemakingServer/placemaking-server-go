package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
)

func RegisterUser(createUserData models.CreateUser) (*models.SanitizedUser, string, error) {
	// Verifica se os e-mails conferem
	if createUserData.Email != createUserData.ConfirmEmail {
		log.Println("error: Emails não conferem")
		return nil, "", fmt.Errorf("emails não conferem")
	}

	// Gera senha temporária e a criptografa
	temporary_password := GenerateUserPassword()

	hashed_password, err := GenerateHashedPassword(temporary_password)
	if err != nil {
		log.Println("erro ao criptografar a senha:", err)
		return nil, "", err
	}

	// Preenche os dados para inserção
	insertUserData := models.InsertUser{
		Name:     createUserData.Name,
		Email:    createUserData.Email,
		Password: hashed_password,
		Role:     createUserData.Role,
		Status:   "active",
	}

	// Insere usuário no banco de dados
	new_user, err := repository.InsertUser(insertUserData)
	if err != nil {
		log.Println("error: Erro ao cadastrar o usuário:", err)
		return nil, "", err
	}

	// Desreferencia o ponteiro para passar um valor, não um ponteiro
	sanitizedUser := models.SanitizeUser(*new_user)
	return &sanitizedUser, temporary_password, nil
}

// LoginUser autentica um usuário e retorna um token JWT se for bem-sucedido.
func LoginUser(email string, password string) (map[string]interface{}, error) {
	// Busca o usuário pelo email
	existingUser, err := repository.GetUserByEmail(email)
	if err != nil {
		return map[string]interface{}{"error": "Usuário não encontrado"}, nil
	}

	// Verifica a senha
	if !CheckPassword(password, existingUser.Password) {
		return map[string]interface{}{"error": "Senha inválida"}, nil
	}

	// Gera um token JWT para o usuário autenticado
	token, err := GenerateUserToken(existingUser, "Bearer")
	if err != nil {
		log.Println("Erro ao gerar token:", err)
		return map[string]interface{}{"error": "Erro ao gerar token"}, nil
	}

	// Retorna os dados sanitizados do usuário e o token
	return map[string]interface{}{
		"token": models.SanitizeToken(*token),
		"user":  models.SanitizeUser(existingUser),
	}, nil
}

func LogoutUser(token string) {
	err := RevokeToken(token)
	if err != nil {
		log.Println("error: erro ao fazer logout:", err)
		return
	}
}

// SendUserData envia um e-mail de boas-vindas para o usuário
func SendUserData(data map[string]interface{}) error {
	// Obtendo os dados do usuário
	userData, ok := data["user"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("erro: estrutura de usuário inválida")
	}

	email, emailOk := userData["email"].(string)
	name, nameOk := userData["name"].(string)
	tempPassword, passOk := data["temporary_password"].(string)

	// Valida se os dados essenciais estão presentes
	if !emailOk || !nameOk || !passOk {
		return fmt.Errorf("erro: dados do usuário incompletos")
	}

	// Caminho do arquivo de template HTML
	templateFilePath := "./assets/templates/welcome.html"

	// Lendo o arquivo do template
	fileContent, err := os.ReadFile(templateFilePath)
	if err != nil {
		return fmt.Errorf("erro ao ler o template de e-mail: %w", err)
	}

	// Criando o template e registrando funções extras
	tmpl, err := template.New("email").Funcs(template.FuncMap{}).Parse(string(fileContent))
	if err != nil {
		return fmt.Errorf("erro ao processar template: %w", err)
	}

	// Criando o conteúdo do e-mail
	renderedEmailContent := new(bytes.Buffer)
	emailData := models.EmailData{
		Name:         name,
		TempPassword: tempPassword,
	}

	// Executando o template com os dados
	err = tmpl.Execute(renderedEmailContent, emailData)
	if err != nil {
		return fmt.Errorf("erro ao renderizar o template: %w", err)
	}

	// Enviar o e-mail
	err = SendEmail(email, renderedEmailContent.String(), "Cadastro realizado com sucesso!")
	if err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	return nil
}

func SendForgotEmailPasswordData(data map[string]interface{}) error {
	fmt.Println("Iniciando envio de e-mail de recuperação...")

	userStruct, ok := data["user"].(models.User)
	if !ok {
		return fmt.Errorf("erro: estrutura de usuário inválida")
	}

	// Serializa a struct para JSON
	userJSON, err := json.Marshal(userStruct)
	if err != nil {
		return fmt.Errorf("erro ao serializar usuário: %w", err)
	}

	// Converte o JSON para um map[string]interface{}
	var userData map[string]interface{}
	err = json.Unmarshal(userJSON, &userData)
	if err != nil {
		return fmt.Errorf("erro ao desserializar usuário: %w", err)
	}
	
	email, emailOk := userData["email"].(string)
	name, nameOk := userData["name"].(string)
	token, tokenOk := data["token"].(string)

	// Valida se os dados essenciais estão presentes
	if !emailOk || !nameOk || !tokenOk {
		fmt.Println("Erro: dados do usuário incompletos")
		return fmt.Errorf("erro: dados do usuário incompletos")
	}

	fmt.Printf("Enviando e-mail para %s (%s), token: %s\n", name, email, token)

	// Caminho do arquivo de template HTML
	templateFilePath := "./assets/templates/recovery-password.html"

	// Lendo o arquivo do template
	fileContent, err := os.ReadFile(templateFilePath)
	if err != nil {
		fmt.Println("Erro ao ler template:", err)
		return fmt.Errorf("erro ao ler o template de e-mail: %w", err)
	}

	// Criando o template e registrando funções extras
	tmpl, err := template.New("email").Funcs(template.FuncMap{}).Parse(string(fileContent))
	if err != nil {
		fmt.Println("Erro ao processar template:", err)
		return fmt.Errorf("erro ao processar template: %w", err)
	}

	// Criando o conteúdo do e-mail
	renderedEmailContent := new(bytes.Buffer)
	emailData := models.RecoveryEmailData{
		Name:  name,
		Token: token,
	}

	// Executando o template com os dados
	err = tmpl.Execute(renderedEmailContent, emailData)
	if err != nil {
		fmt.Println("Erro ao renderizar template:", err)
		return fmt.Errorf("erro ao renderizar o template: %w", err)
	}

	fmt.Println("Conteúdo do e-mail gerado com sucesso!")

	// Enviar o e-mail
	err = SendEmail(email, renderedEmailContent.String(), "Recuperação de Senha")
	if err != nil {
		fmt.Println("Erro ao enviar e-mail:", err)
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	fmt.Println("E-mail enviado com sucesso!")
	return nil
}

func ForgotPasswordService(email string) (map[string]interface{}, error) {

	user, err := repository.GetUserByEmail(email)

	if err != nil {
		return map[string]interface{}{"error": "Usuário não encontrado"}, nil
	}

	token, err := GenerateUserToken(user, "RecoverPassword")

	if err != nil {
		return map[string]interface{}{"error": "Erro ao gerar token"}, nil
	}

	return map[string]interface{}{"error": nil, "token": token.Token, "user": user}, nil
}

func ValidateCodeService(token string) (map[string]interface{}, error) {
	tokenToValidate, err := repository.GetToken(token)

	if err != nil {
		return map[string]interface{}{"error": "Código inválido ou não encontrado."}, err
	}

	if tokenToValidate.Token != token {
		return map[string]interface{}{"error": "Código inválido."}, nil
	}

	user, err := repository.GetUserById(tokenToValidate.User_id)

	if err != nil {
		return map[string]interface{}{"error": "Erro ao buscar usuário no banco de dados."}, err
	}

	accessToken, err := GenerateUserToken(user, "Bearer")

	if err != nil {
		return map[string]interface{}{"error": "Erro ao gerar token de usuário"}, err
	}

	return map[string]interface{}{
		"error": nil, 
		"access_token": models.SanitizeToken(*accessToken), 
		"validated_token": tokenToValidate.Token, 
		"user": user}, nil
}

func ResetPasswordService(newPassword, confirmPassword, userID string) (map[string]interface{}, error) {

	if newPassword != confirmPassword {
		return map[string]interface{}{"error": "Senhas não conferem"}, nil
	}

	hashedPassword, err := GenerateHashedPassword(newPassword)
	if err != nil {
		return map[string]interface{}{"error": "Erro ao criptografar a nova senha."}, err
	}

	err = repository.UpdateUserPassword(userID, hashedPassword)
	if err != nil {
		return map[string]interface{}{"error": "Erro ao atualizar senha no banco."}, err
	}

	return map[string]interface{}{"message": "Senha redefinida com sucesso!"}, nil
}
