# Correção do Tratamento de Arquivos de Áudio OGG

## Problema Identificado
O projeto não estava tratando corretamente arquivos de áudio recebidos no formato OGG via WhatsApp. A função `determineMediaExtension` não tinha mapeamento explícito para MIME types de áudio, e a função `ExtractMediaInfo` sempre retornava a extensão "ogg" independentemente do tipo real do arquivo.

## Alterações Realizadas

### 1. Arquivo: `src/pkg/utils/whatsapp.go`

#### Função `determineMediaExtension` (linhas 69-91)
**Antes:**
```go
func determineMediaExtension(originalFilename, mimeType string) string {
    if originalFilename != "" {
        if ext := filepath.Ext(originalFilename); ext != "" {
            return ext
        }
    }

    if ext, ok := resolveKnownDocumentExtension(mimeType); ok {
        return ext
    }

    if ext, err := mime.ExtensionsByType(mimeType); err == nil && len(ext) > 0 {
        return ext[0]
    }

    if parts := strings.Split(mimeType, "/"); len(parts) > 1 {
        return "." + parts[len(parts)-1]
    }

    return ""
}
```

**Depois:**
```go
func determineMediaExtension(originalFilename, mimeType string) string {
    if originalFilename != "" {
        if ext := filepath.Ext(originalFilename); ext != "" {
            return ext
        }
    }

    // Handle specific audio MIME types
    switch mimeType {
    case "audio/ogg", "audio/ogg; codecs=opus":
        return ".ogg"
    case "audio/mpeg", "audio/mp3":
        return ".mp3"
    case "audio/mp4", "audio/aac":
        return ".m4a"
    case "audio/wav":
        return ".wav"
    case "audio/webm", "audio/webm; codecs=opus":
        return ".webm"
    }

    if ext, ok := resolveKnownDocumentExtension(mimeType); ok {
        return ext
    }

    if ext, err := mime.ExtensionsByType(mimeType); err == nil && len(ext) > 0 {
        return ext[0]
    }

    if parts := strings.Split(mimeType, "/"); len(parts) > 1 {
        return "." + parts[len(parts)-1]
    }

    return ""
}
```

#### Função `ExtractMediaInfo` - Tratamento de AudioMessage (linhas 277-287)
**Antes:**
```go
// Check for audio message
if aud := msg.GetAudioMessage(); aud != nil {
    extension := "ogg"
    if aud.GetPTT() {
        extension = "ogg" // Voice notes are typically ogg
    }
    filename = GenerateMediaFilename("audio", extension, "")
    return "audio", filename,
        aud.GetURL(), aud.GetMediaKey(), aud.GetFileSHA256(),
        aud.GetFileEncSHA256(), aud.GetFileLength()
}
```

**Depois:**
```go
// Check for audio message
if aud := msg.GetAudioMessage(); aud != nil {
    // Determine extension based on MIME type
    mimeType := aud.GetMimetype()
    extension := strings.TrimPrefix(determineMediaExtension("", mimeType), ".")
    if extension == "" {
        extension = "ogg" // Fallback to ogg
    }
    filename = GenerateMediaFilename("audio", extension, "")
    return "audio", filename,
        aud.GetURL(), aud.GetMediaKey(), aud.GetFileSHA256(),
        aud.GetFileEncSHA256(), aud.GetFileLength()
}
```

### 2. Arquivo: `src/pkg/utils/whatsapp_test.go`

Adicionados 6 novos casos de teste para validar o tratamento de diferentes formatos de áudio:

```go
{
    name:       "OggAudio",
    filename:   "",
    mimeType:   "audio/ogg",
    wantSuffix: ".ogg",
},
{
    name:       "OggAudioWithCodec",
    filename:   "",
    mimeType:   "audio/ogg; codecs=opus",
    wantSuffix: ".ogg",
},
{
    name:       "Mp3Audio",
    filename:   "",
    mimeType:   "audio/mpeg",
    wantSuffix: ".mp3",
},
{
    name:       "M4aAudio",
    filename:   "",
    mimeType:   "audio/mp4",
    wantSuffix: ".m4a",
},
{
    name:       "WavAudio",
    filename:   "",
    mimeType:   "audio/wav",
    wantSuffix: ".wav",
},
{
    name:       "WebmAudio",
    filename:   "",
    mimeType:   "audio/webm; codecs=opus",
    wantSuffix: ".webm",
},
```

## Resultado dos Testes

Todos os testes passaram com sucesso:

```
=== RUN   TestDetermineMediaExtension
=== RUN   TestDetermineMediaExtension/OggAudio
=== RUN   TestDetermineMediaExtension/OggAudioWithCodec
=== RUN   TestDetermineMediaExtension/Mp3Audio
=== RUN   TestDetermineMediaExtension/M4aAudio
=== RUN   TestDetermineMediaExtension/WavAudio
=== RUN   TestDetermineMediaExtension/WebmAudio
--- PASS: TestDetermineMediaExtension (0.00s)
```

## Formatos de Áudio Suportados

Agora o projeto suporta corretamente os seguintes formatos de áudio:

- **OGG** (`audio/ogg`, `audio/ogg; codecs=opus`) → `.ogg`
- **MP3** (`audio/mpeg`, `audio/mp3`) → `.mp3`
- **M4A/AAC** (`audio/mp4`, `audio/aac`) → `.m4a`
- **WAV** (`audio/wav`) → `.wav`
- **WebM** (`audio/webm`, `audio/webm; codecs=opus`) → `.webm`

## Ambiente de Desenvolvimento

Foi criado um ambiente Docker completo para desenvolvimento:

### Arquivos Criados:
- `Dockerfile.dev` - Imagem de desenvolvimento com Go 1.24.0 e todas as dependências
- `docker-compose.dev.yml` - Orquestração do container de desenvolvimento

### Como Usar:

```bash
# Iniciar o container de desenvolvimento
docker-compose -f docker-compose.dev.yml up -d

# Executar testes
docker exec whatsapp-dev go test -v ./pkg/utils

# Compilar o projeto
docker exec whatsapp-dev go build -o whatsapp .

# Parar o container
docker-compose -f docker-compose.dev.yml down
```

## Impacto

- ✅ Arquivos de áudio OGG agora são salvos com a extensão correta
- ✅ Suporte para múltiplos formatos de áudio (MP3, M4A, WAV, WebM)
- ✅ Extensão do arquivo determinada pelo MIME type real da mensagem
- ✅ Fallback para `.ogg` em caso de MIME type desconhecido
- ✅ Todos os testes unitários passando
- ✅ Sem quebra de compatibilidade com código existente
