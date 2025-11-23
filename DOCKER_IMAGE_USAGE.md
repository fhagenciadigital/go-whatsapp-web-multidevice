# 🎉 Imagem Docker Publicada com Sucesso!

## 📦 Detalhes da Imagem

- **Docker Hub:** https://hub.docker.com/r/fhagenciadigital/go-whatsapp-web-multidevice
- **Nome:** `fhagenciadigital/go-whatsapp-web-multidevice`
- **Tags disponíveis:**
  - `latest` - Versão mais recente
  - `audio-fix` - Versão com correção de áudio OGG
- **Tamanho:** ~62 MB (comprimido)
- **Arquitetura:** linux/amd64

## ✨ O que está incluído

Esta imagem contém:
- ✅ Correção para arquivos de áudio OGG
- ✅ Suporte a múltiplos formatos de áudio (OGG, MP3, M4A, WAV, WebM)
- ✅ Go 1.24.0
- ✅ FFmpeg para processamento de mídia
- ✅ Alpine Linux (otimizado para tamanho reduzido)

## 🚀 Como Usar

### Método 1: Docker Run

```bash
docker run -d \
  --name whatsapp \
  -p 3000:3000 \
  -v whatsapp-data:/app/storages \
  fhagenciadigital/go-whatsapp-web-multidevice:latest
```

Acesse: http://localhost:3000

### Método 2: Docker Compose (Recomendado)

```bash
docker-compose -f docker-compose.custom.yml up -d
```

### Método 3: Com Configurações Personalizadas

```bash
docker run -d \
  --name whatsapp \
  -p 3000:3000 \
  -v whatsapp-data:/app/storages \
  -e APP_DEBUG=true \
  -e APP_BASIC_AUTH=admin:senha123 \
  -e WHATSAPP_AUTO_MARK_READ=true \
  fhagenciadigital/go-whatsapp-web-multidevice:latest
```

## 🔧 Variáveis de Ambiente

| Variável | Descrição | Padrão |
|----------|-----------|---------|
| `APP_PORT` | Porta da aplicação | `3000` |
| `APP_DEBUG` | Modo debug | `false` |
| `APP_OS` | Nome do dispositivo | `Chrome` |
| `APP_BASIC_AUTH` | Autenticação básica | - |
| `WHATSAPP_AUTO_REPLY` | Resposta automática | - |
| `WHATSAPP_AUTO_MARK_READ` | Marcar como lido | `false` |
| `WHATSAPP_WEBHOOK` | URL do webhook | - |

## 📝 Exemplos de Uso

### REST API Mode (padrão)

```bash
docker run -d \
  --name whatsapp-rest \
  -p 3000:3000 \
  fhagenciadigital/go-whatsapp-web-multidevice:latest rest
```

### MCP Server Mode

```bash
docker run -d \
  --name whatsapp-mcp \
  -p 8080:8080 \
  fhagenciadigital/go-whatsapp-web-multidevice:latest mcp
```

### Com Webhook

```bash
docker run -d \
  --name whatsapp-webhook \
  -p 3000:3000 \
  -e WHATSAPP_WEBHOOK=https://seu-webhook.com/handler \
  -e WHATSAPP_WEBHOOK_SECRET=seu-segredo \
  fhagenciadigital/go-whatsapp-web-multidevice:latest
```

## 🔄 Atualizar para Nova Versão

```bash
# Parar o container atual
docker stop whatsapp
docker rm whatsapp

# Baixar nova versão
docker pull fhagenciadigital/go-whatsapp-web-multidevice:latest

# Iniciar com a nova versão
docker run -d \
  --name whatsapp \
  -p 3000:3000 \
  -v whatsapp-data:/app/storages \
  fhagenciadigital/go-whatsapp-web-multidevice:latest
```

## 🐛 Verificar Logs

```bash
# Ver logs em tempo real
docker logs -f whatsapp

# Ver últimas 100 linhas
docker logs --tail 100 whatsapp
```

## 💾 Backup dos Dados

```bash
# Criar backup do volume
docker run --rm -v whatsapp-data:/data -v ${PWD}:/backup \
  alpine tar czf /backup/whatsapp-backup.tar.gz /data

# Restaurar backup
docker run --rm -v whatsapp-data:/data -v ${PWD}:/backup \
  alpine tar xzf /backup/whatsapp-backup.tar.gz -C /
```

## 🔗 Links Úteis

- **Docker Hub:** https://hub.docker.com/r/fhagenciadigital/go-whatsapp-web-multidevice
- **Documentação API:** http://localhost:3000 (após iniciar)
- **Projeto Original:** https://github.com/aldinokemal/go-whatsapp-web-multidevice

## 📊 Estatísticas da Imagem

```bash
# Ver informações da imagem
docker inspect fhagenciadigital/go-whatsapp-web-multidevice:latest

# Ver tamanho
docker images fhagenciadigital/go-whatsapp-web-multidevice

# Ver histórico de camadas
docker history fhagenciadigital/go-whatsapp-web-multidevice:latest
```

## 🆘 Suporte

Se encontrar algum problema:

1. Verifique os logs: `docker logs whatsapp`
2. Verifique se a porta 3000 está livre: `netstat -an | findstr :3000`
3. Verifique se o Docker está em execução: `docker ps`

---

**Última atualização:** 23 de Novembro de 2025
**Versão:** 7.9.0 com correção de áudio OGG
