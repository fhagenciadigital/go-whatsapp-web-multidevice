# 🎉 Nova Versão Publicada no Docker Hub

**Data:** 23 de Novembro de 2025  
**Versão:** v7.9.1

## 📦 Imagens Publicadas

**Docker Hub:** https://hub.docker.com/r/fhagenciadigital/go-whatsapp-web-multidevice

### Tags Disponíveis:

| Tag | Descrição | Image ID | Tamanho |
|-----|-----------|----------|---------|
| `latest` | Versão mais recente (v7.9.1) | `4ea15c32ea5c` | 62.2 MB |
| `v7.9.1` | Versão específica com todas as features | `4ea15c32ea5c` | 62.2 MB |
| `media-cleanup` | Tag específica para feature de limpeza | `4ea15c32ea5c` | 62.2 MB |
| `audio-fix` | Versão anterior (apenas correção OGG) | `5ef70588a789` | 62.2 MB |

## ✨ Novas Funcionalidades (v7.9.1)

### 1. ✅ Correção de Áudio OGG
- Tratamento correto de arquivos de áudio OGG recebidos via WhatsApp
- Suporte para múltiplos formatos: OGG, MP3, M4A, WAV, WebM
- Extensão de arquivo determinada pelo MIME type real

### 2. 🧹 Limpeza Automática de Mídia (NOVA!)
- Remove automaticamente arquivos de mídia antigos
- Configurável via variáveis de ambiente
- Executa em background sem impactar performance
- Logs detalhados de operações
- Remove diretórios vazios automaticamente

## 🚀 Como Usar a Nova Versão

### Opção 1: Docker Run Simples

```bash
docker pull fhagenciadigital/go-whatsapp-web-multidevice:latest

docker run -d \
  --name whatsapp \
  -p 3000:3000 \
  -v whatsapp-data:/app/storages \
  fhagenciadigital/go-whatsapp-web-multidevice:latest
```

### Opção 2: Com Limpeza Automática Habilitada

```bash
docker run -d \
  --name whatsapp \
  -p 3000:3000 \
  -v whatsapp-data:/app/storages \
  -e MEDIA_CLEANUP_ENABLED=true \
  -e MEDIA_CLEANUP_RETENTION_DAYS=30 \
  -e MEDIA_CLEANUP_INTERVAL_HOURS=24 \
  fhagenciadigital/go-whatsapp-web-multidevice:latest
```

### Opção 3: Docker Compose

```yaml
services:
  whatsapp:
    image: fhagenciadigital/go-whatsapp-web-multidevice:latest
    container_name: whatsapp
    restart: always
    ports:
      - "3000:3000"
    volumes:
      - whatsapp-data:/app/storages
    environment:
      - APP_PORT=3000
      - APP_DEBUG=true
      - APP_OS=Chrome
      # Limpeza automática de mídia
      - MEDIA_CLEANUP_ENABLED=true
      - MEDIA_CLEANUP_RETENTION_DAYS=30
      - MEDIA_CLEANUP_INTERVAL_HOURS=24

volumes:
  whatsapp-data:
```

## 📋 Variáveis de Ambiente Disponíveis

### Aplicação
- `APP_PORT` - Porta da aplicação (padrão: 3000)
- `APP_DEBUG` - Modo debug (padrão: false)
- `APP_OS` - Nome do dispositivo (padrão: Chrome)
- `APP_BASIC_AUTH` - Autenticação básica

### WhatsApp
- `WHATSAPP_AUTO_REPLY` - Resposta automática
- `WHATSAPP_AUTO_MARK_READ` - Marcar como lido (padrão: false)
- `WHATSAPP_AUTO_DOWNLOAD_MEDIA` - Download automático (padrão: true)
- `WHATSAPP_WEBHOOK` - URL do webhook

### Limpeza de Mídia (NOVO!)
- `MEDIA_CLEANUP_ENABLED` - Habilita limpeza (padrão: false)
- `MEDIA_CLEANUP_RETENTION_DAYS` - Dias de retenção (padrão: 7)
- `MEDIA_CLEANUP_INTERVAL_HOURS` - Intervalo de execução (padrão: 24)

## 🧪 Teste da Publicação

✅ **Build:** Sucesso (358 segundos)  
✅ **Push latest:** Sucesso  
✅ **Push v7.9.1:** Sucesso  
✅ **Push media-cleanup:** Sucesso  
✅ **Teste local:** Sucesso  
✅ **Logs confirmados:** Feature de limpeza ativa

### Log de Confirmação:
```
INFO Starting media cleanup scheduler: retention=7 days, interval=24 hours
```

## 📊 Comparação de Versões

| Feature | audio-fix | v7.9.1 (latest) |
|---------|-----------|-----------------|
| Correção áudio OGG | ✅ | ✅ |
| Suporte múltiplos formatos | ✅ | ✅ |
| Limpeza automática de mídia | ❌ | ✅ |
| Configuração via env vars | Parcial | Completa |
| Testes unitários | ✅ | ✅ |
| Documentação | Básica | Completa |

## 🔄 Atualizar de Versão Anterior

```bash
# 1. Parar container atual
docker stop whatsapp
docker rm whatsapp

# 2. Baixar nova versão
docker pull fhagenciadigital/go-whatsapp-web-multidevice:latest

# 3. Iniciar com nova versão
docker run -d \
  --name whatsapp \
  -p 3000:3000 \
  -v whatsapp-data:/app/storages \
  -e MEDIA_CLEANUP_ENABLED=true \
  -e MEDIA_CLEANUP_RETENTION_DAYS=30 \
  fhagenciadigital/go-whatsapp-web-multidevice:latest

# 4. Verificar logs
docker logs -f whatsapp
```

## 📝 Verificar Versão Instalada

```bash
# Ver tags disponíveis localmente
docker images fhagenciadigital/go-whatsapp-web-multidevice

# Ver logs de inicialização
docker logs whatsapp 2>&1 | grep -i "cleanup\|fiber"

# Verificar no Docker Hub
# https://hub.docker.com/r/fhagenciadigital/go-whatsapp-web-multidevice/tags
```

## 🆘 Suporte e Documentação

### Documentação Completa:
- `MEDIA_CLEANUP_FEATURE.md` - Guia completo da feature de limpeza
- `DOCKER_IMAGE_USAGE.md` - Como usar a imagem Docker
- `readme.md` - Documentação geral do projeto

### Exemplos de Configuração:
- `docker-compose.custom.yml` - Exemplo Docker Compose
- `src/.env.example` - Exemplo de variáveis de ambiente

## 🎯 Cenários Recomendados

### Desenvolvimento (7 dias)
```bash
MEDIA_CLEANUP_ENABLED=true
MEDIA_CLEANUP_RETENTION_DAYS=7
MEDIA_CLEANUP_INTERVAL_HOURS=24
```

### Produção - Uso Moderado (30 dias)
```bash
MEDIA_CLEANUP_ENABLED=true
MEDIA_CLEANUP_RETENTION_DAYS=30
MEDIA_CLEANUP_INTERVAL_HOURS=24
```

### Produção - Alto Volume (7 dias, cleanup frequente)
```bash
MEDIA_CLEANUP_ENABLED=true
MEDIA_CLEANUP_RETENTION_DAYS=7
MEDIA_CLEANUP_INTERVAL_HOURS=12
```

### Arquivamento (90+ dias)
```bash
MEDIA_CLEANUP_ENABLED=true
MEDIA_CLEANUP_RETENTION_DAYS=90
MEDIA_CLEANUP_INTERVAL_HOURS=168
```

## ⚠️ Notas Importantes

1. **Backup:** Sempre faça backup antes de habilitar a limpeza automática
2. **Teste:** Teste em ambiente de desenvolvimento primeiro
3. **Monitoramento:** Verifique os logs regularmente
4. **Espaço:** Monitore o espaço em disco disponível
5. **Padrão:** A limpeza está DESABILITADA por padrão (seguro)

## 📈 Estatísticas da Build

- **Tempo de build:** 358 segundos (~6 minutos)
- **Tamanho final:** 62.2 MB (comprimido)
- **Tamanho descomprimido:** 229 MB
- **Layers:** 5
- **Arquitetura:** linux/amd64

## ✅ Checklist de Publicação

- [x] Código implementado e testado
- [x] Build da imagem Docker concluído
- [x] Tag `latest` publicada
- [x] Tag `v7.9.1` publicada
- [x] Tag `media-cleanup` publicada
- [x] Teste local realizado
- [x] Logs verificados
- [x] Documentação atualizada
- [x] Docker Hub atualizado

## 🔗 Links

- **Docker Hub:** https://hub.docker.com/r/fhagenciadigital/go-whatsapp-web-multidevice
- **GitHub:** https://github.com/fhagenciadigital/go-whatsapp-web-multidevice
- **Projeto Original:** https://github.com/aldinokemal/go-whatsapp-web-multidevice

---

**Status:** ✅ Publicado e Testado  
**Versão:** v7.9.1  
**Data:** 23 de Novembro de 2025  
**Publicado por:** GitHub Copilot
