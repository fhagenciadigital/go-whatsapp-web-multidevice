# Teste de Envio de Áudio Binário

## Alterações Implementadas

### 1. Validação Melhorada (`send_validation.go`)
- **Antes**: Rejeitava ficheiros com Content-Type vazio ou `application/octet-stream`
- **Agora**: Aceita Content-Type vazio ou genérico, deixando a detecção para o processamento

### 2. Nova Função de Detecção (`whatsapp.go`)
Adicionada função `DetectAudioMimeType()` que:
- Detecta MIME type pela extensão do ficheiro
- Verifica "magic bytes" (assinaturas de ficheiro):
  - OGG: `4F 67 67 53` ("OggS")
  - MP3: `ID3` ou `FF FB/FF F3/FF F2`
  - M4A/AAC: `ftyp` no offset 4
  - WAV: `RIFF....WAVE`
  - FLAC: `fLaC`
  - AMR: `#!AMR`

### 3. Processamento Aprimorado (`send.go`)
- Primeiro tenta `http.DetectContentType()` do Go
- Se falhar (retornar octet-stream), usa `DetectAudioMimeType()`
- Adiciona logging para debugging:
  - MIME type detectado
  - Nome do ficheiro
  - Tamanho em bytes
  - Informação de upload

### 4. Validação de Conteúdo
- Verifica se o ficheiro não está vazio
- Rejeita uploads com 0 bytes

## Como Testar

### Exemplo com cURL (Windows PowerShell)

```powershell
# Preparar o ficheiro de teste
$audioFile = "C:\path\to\your\audio.ogg"

# Enviar via cURL
curl.exe -X POST "http://localhost:3000/send/audio" `
  -H "Content-Type: multipart/form-data" `
  -F "phone=351912345678@s.whatsapp.net" `
  -F "audio=@$audioFile"
```

### Exemplo com n8n

**HTTP Request Node:**
- Method: `POST`
- URL: `http://your-server:3000/send/audio`
- Body Content Type: `Form-Data Multipart`
- Body Parameters:
  - `phone`: `351912345678@s.whatsapp.net`
  - `audio`: *Selecionar ficheiro binário*

### Exemplo com Python

```python
import requests

url = "http://localhost:3000/send/audio"
files = {
    'audio': ('audio.ogg', open('audio.ogg', 'rb'), 'audio/ogg')
}
data = {
    'phone': '351912345678@s.whatsapp.net'
}

response = requests.post(url, files=files, data=data)
print(response.json())
```

## Formatos de Áudio Suportados

| Formato | MIME Types | Extensões |
|---------|-----------|-----------|
| OGG | audio/ogg | .ogg |
| MP3 | audio/mpeg, audio/mp3 | .mp3 |
| M4A/AAC | audio/mp4, audio/aac, audio/m4a | .m4a, .aac |
| WAV | audio/wav, audio/vnd.wav, audio/wave | .wav |
| FLAC | audio/flac | .flac |
| AMR | audio/amr | .amr |
| WMA | audio/wma, audio/x-ms-wma | .wma |
| WebM | audio/webm | .webm |

## Verificar Logs

Após enviar o áudio, verificar os logs do container:

```powershell
docker logs <container-name> --tail 50
```

Procurar por linhas como:
```
INFO[...] Audio MIME type detected from filename/signature: audio/ogg (filename: audio.ogg, size: 245123 bytes)
INFO[...] Uploading audio to WhatsApp: MIME=audio/ogg, size=245123 bytes
```

## Troubleshooting

### Se o áudio ainda não aparecer:

1. **Verificar formato do telefone**:
   - Deve incluir código do país
   - Deve terminar com `@s.whatsapp.net`
   - Exemplo: `351912345678@s.whatsapp.net`

2. **Verificar tamanho do ficheiro**:
   - WhatsApp tem limites de tamanho
   - Para ficheiros grandes, usar `audio_url` em vez de upload direto

3. **Verificar se o container está a correr**:
   ```powershell
   docker ps
   ```

4. **Verificar logs para erros**:
   ```powershell
   docker logs <container-name> --tail 100 | Select-String "error|ERROR|fail|FAIL"
   ```

5. **Testar com formato MP3 simples**:
   - MP3 é o formato mais amplamente suportado
   - Converter o áudio para MP3 e testar novamente

## Reconstruir a Aplicação

Para aplicar as alterações:

```powershell
# Parar o container atual
docker stop <container-name>
docker rm <container-name>

# Reconstruir a imagem
docker build -f docker/golang.Dockerfile -t go-whatsapp:test .

# Executar novo container
docker run -d `
  -p 3000:3000 `
  -v ${PWD}/storages:/app/storages `
  --name whatsapp-test `
  go-whatsapp:test
```

Ou usar o Docker Compose:

```powershell
docker-compose -f docker-compose.custom.yml down
docker-compose -f docker-compose.custom.yml build
docker-compose -f docker-compose.custom.yml up -d
```
