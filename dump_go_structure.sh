#!/usr/bin/env bash
# ---------------------------------------------------------------------------
# dump_go_structure.sh – Exporta estrutura + conteúdo de uma aplicação Go
# ---------------------------------------------------------------------------

set -euo pipefail

ROOT_DIR="$(pwd)"
TIMESTAMP="$(date +'%Y-%m-%d_%H-%M-%S')"
OUTPUT_FILE="${ROOT_DIR}/go_structure_dump_${TIMESTAMP}.txt"

# extensões relevantes para projetos Go
EXTENSIONS=(
  "*.go"
  "*.mod" "*.sum"
  "*.yml" "*.yaml"
  "*.json" "*.env"
  "*.sql"
  "*.sh"
  "*.md"
)

# pastas a excluir do dump
EXCLUDE_DIRS=(
  ".git" "vendor" "bin" "dist" "coverage"
  "node_modules" ".cache" "testdata" "tmp"
)

# cria arquivo de saída
mkdir -p "$(dirname "$OUTPUT_FILE")"
: > "$OUTPUT_FILE"

cat <<EOF >> "$OUTPUT_FILE"
🔎 Dump de Estrutura – Projeto Go
Raiz do projeto : $ROOT_DIR
Gerado em       : $(date)
------------------------------------------------------------

EOF

# monta cláusula de exclusão
EXCLUDE_CLAUSE=()
for dir in "${EXCLUDE_DIRS[@]}"; do
  EXCLUDE_CLAUSE+=( -name "$dir" -o )
done
unset 'EXCLUDE_CLAUSE[${#EXCLUDE_CLAUSE[@]}-1]'

# monta cláusula de extensão
EXTENSION_CLAUSE=()
for ext in "${EXTENSIONS[@]}"; do
  EXTENSION_CLAUSE+=( -iname "$ext" -o )
done
unset 'EXTENSION_CLAUSE[${#EXTENSION_CLAUSE[@]}-1]'

# executa o find e escreve no dump
find "$ROOT_DIR" \
  \( -type d \( "${EXCLUDE_CLAUSE[@]}" \) -prune \) -o \
  -type f \( "${EXTENSION_CLAUSE[@]}" \) -print0 |
sort -z |
while IFS= read -r -d '' file; do
  filename="$(basename "$file")"
  relative_path="${file#"${ROOT_DIR}/"}"
  folder="$(dirname "$relative_path")"

  {
    echo "📄 Arquivo : $filename"
    echo "📂 Pasta   : $folder"
    echo "🧭 Caminho : $relative_path"
    echo "--------------------------------------"
    echo "📜 Conteúdo:"
    echo
  } >> "$OUTPUT_FILE"

  if file "$file" | grep -qE 'image|binary|ELF|compressed'; then
    echo "[Arquivo binário / imagem – conteúdo omitido]" >> "$OUTPUT_FILE"
  else
    cat "$file" >> "$OUTPUT_FILE"
  fi

  echo -e "\n\n============================================================\n\n" >> "$OUTPUT_FILE"
done

echo "✅ Estrutura exportada em: $OUTPUT_FILE"
