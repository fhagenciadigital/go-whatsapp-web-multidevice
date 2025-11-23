# 🧹 Limpeza Automática de Arquivos de Mídia

## 📋 Visão Geral

Foi adicionada uma funcionalidade de limpeza automática de arquivos de mídia antigos. Esta funcionalidade executa em background e remove automaticamente arquivos de mídia (imagens, vídeos, áudio, documentos, etc.) que ultrapassem um determinado número de dias de retenção.

## ✨ Funcionalidades

- ✅ Limpeza automática programada de arquivos de mídia
- ✅ Configurável via variáveis de ambiente ou flags de linha de comando
- ✅ Remoção de diretórios vazios após limpeza
- ✅ Logs detalhados de operações de limpeza
- ✅ Execução inicial após 30 segundos do startup
- ✅ Execução periódica configurável (padrão: 24 horas)
- ✅ Estatísticas de limpeza (arquivos deletados, tamanho liberado)

## 🔧 Configuração

### Variáveis de Ambiente

| Variável | Descrição | Padrão | Exemplo |
|----------|-----------|--------|---------|
| `MEDIA_CLEANUP_ENABLED` | Habilita limpeza automática | `false` | `true` |
| `MEDIA_CLEANUP_RETENTION_DAYS` | Dias para manter arquivos (0=desabilitado) | `7` | `30` |
| `MEDIA_CLEANUP_INTERVAL_HOURS` | Horas entre execuções de limpeza | `24` | `12` |

### Flags de Linha de Comando

```bash
# Habilitar limpeza com 30 dias de retenção
./whatsapp rest --media-cleanup-enabled=true --media-cleanup-retention-days=30

# Executar limpeza a cada 12 horas
./whatsapp rest --media-cleanup-enabled=true --media-cleanup-interval-hours=12
```

## 🐳 Uso com Docker

### Docker Run

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

### Docker Compose

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
      - MEDIA_CLEANUP_ENABLED=true
      - MEDIA_CLEANUP_RETENTION_DAYS=30    # Manter arquivos por 30 dias
      - MEDIA_CLEANUP_INTERVAL_HOURS=24    # Executar limpeza a cada 24 horas

volumes:
  whatsapp-data:
```

## 📊 Comportamento

### Quando a Limpeza é Executada?

1. **Primeira execução:** 30 segundos após o startup da aplicação
2. **Execuções subsequentes:** A cada N horas (configurável via `MEDIA_CLEANUP_INTERVAL_HOURS`)

### O que é Limpo?

- Arquivos na pasta `statics/media/` e subpastas
- Apenas arquivos com data de modificação anterior a X dias (configurável)
- Diretórios vazios são removidos após limpeza de arquivos

### O que NÃO é Limpo?

- Arquivos com menos de X dias
- Arquivos em outras pastas (`statics/qrcode`, `statics/senditems`, `storages`)
- Database e configurações

## 📝 Exemplos de Logs

### Limpeza Habilitada com Sucesso

```
INFO Starting media cleanup scheduler: retention=30 days, interval=24 hours
INFO Starting media cleanup: removing files older than 30 days (before 2025-10-24 12:00:00)
INFO Media cleanup completed: scanned 150 files (250.50 MB), deleted 45 files (120.30 MB), failed 0
```

### Limpeza Desabilitada

```
INFO Media cleanup scheduler disabled (retention days = 0)
```

### Arquivo Individual Deletado (Debug Mode)

```
DEBUG Deleted old media file: statics/media/628123456789/audio_1729512000_abc123.ogg (age: 35 days, size: 2048576 bytes)
DEBUG Removed empty directory: statics/media/628123456789
```

## 🎯 Casos de Uso

### Desenvolvimento / Testes (7 dias)

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

### Produção - Alto Volume (7 dias, limpeza frequente)

```bash
MEDIA_CLEANUP_ENABLED=true
MEDIA_CLEANUP_RETENTION_DAYS=7
MEDIA_CLEANUP_INTERVAL_HOURS=12
```

### Arquivamento (90 dias ou mais)

```bash
MEDIA_CLEANUP_ENABLED=true
MEDIA_CLEANUP_RETENTION_DAYS=90
MEDIA_CLEANUP_INTERVAL_HOURS=168  # 1 semana
```

### Desabilitado (manter todos os arquivos)

```bash
MEDIA_CLEANUP_ENABLED=false
# ou
MEDIA_CLEANUP_RETENTION_DAYS=0
```

## ⚠️ Considerações Importantes

### Backup

- **Sempre faça backup dos dados antes de habilitar a limpeza automática**
- Arquivos deletados não podem ser recuperados
- Considere usar volumes Docker para facilitar backups

### Espaço em Disco

- Monitore o espaço em disco disponível
- Ajuste `MEDIA_CLEANUP_RETENTION_DAYS` conforme necessário
- Em ambientes com muito tráfego, considere valores menores (7-15 dias)

### Performance

- A limpeza é executada em background (goroutine)
- Não bloqueia operações normais da aplicação
- Em diretórios muito grandes, pode levar alguns minutos

### Debug

Para ver logs detalhados da limpeza:

```bash
docker run -d \
  -e APP_DEBUG=true \
  -e MEDIA_CLEANUP_ENABLED=true \
  -e MEDIA_CLEANUP_RETENTION_DAYS=30 \
  fhagenciadigital/go-whatsapp-web-multidevice:latest
```

## 🔍 Verificação

### Ver Logs de Limpeza

```bash
# Ver logs em tempo real
docker logs -f whatsapp | grep "cleanup"

# Ver apenas logs de limpeza
docker logs whatsapp 2>&1 | grep -i "media cleanup"
```

### Verificar Tamanho da Pasta de Mídia

```bash
# Dentro do container
docker exec whatsapp du -sh /app/statics/media

# Listar arquivos por tamanho
docker exec whatsapp find /app/statics/media -type f -exec ls -lh {} \; | sort -k5 -hr | head -20
```

### Testar Manualmente

Você pode testar a limpeza manualmente:

```go
// Dentro do código Go
utils.CleanupOldMediaFiles(30) // Remove arquivos com mais de 30 dias
```

## 📦 Arquivos Adicionados/Modificados

### Novos Arquivos

1. **`src/pkg/utils/media_cleanup.go`**
   - Função `CleanupOldMediaFiles()` - Executa limpeza
   - Função `StartMediaCleanupScheduler()` - Inicia scheduler
   - Função `cleanupEmptyDirectories()` - Remove diretórios vazios

2. **`src/pkg/utils/media_cleanup_test.go`**
   - Testes unitários completos
   - Cobertura de múltiplos cenários

### Arquivos Modificados

1. **`src/config/settings.go`**
   - Adicionadas variáveis: `MediaCleanupEnabled`, `MediaCleanupRetentionDays`, `MediaCleanupIntervalHours`

2. **`src/cmd/root.go`**
   - Leitura de variáveis de ambiente
   - Flags de linha de comando
   - Inicialização do scheduler

3. **`readme.md`**
   - Documentação das novas variáveis de ambiente

4. **`docker-compose.custom.yml`**
   - Exemplos de configuração

## 🧪 Testes

Execute os testes:

```bash
# No container de desenvolvimento
docker exec whatsapp-dev go test -v ./pkg/utils -run TestCleanupOldMediaFiles

# Todos os testes do pacote utils
docker exec whatsapp-dev go test -v ./pkg/utils
```

Todos os testes devem passar:

```
✓ TestCleanupOldMediaFiles/Disabled_cleanup
✓ TestCleanupOldMediaFiles/Delete_files_older_than_7_days
✓ TestCleanupOldMediaFiles/Delete_files_older_than_30_days
✓ TestCleanupOldMediaFiles_WithSubdirectories
✓ TestCleanupOldMediaFiles_NonExistentPath
```

## 🚀 Próximos Passos

Para utilizar a funcionalidade:

1. **Rebuildar a imagem Docker:**
   ```bash
   cd c:\Users\paulo\repos\go-whatsapp-web-multidevice
   .\build-and-push.ps1 -DockerHubUsername "fhagenciadigital"
   ```

2. **Ou usar localmente:**
   ```bash
   docker build -f docker/golang.Dockerfile -t whatsapp-local .
   docker run -d -p 3000:3000 \
     -e MEDIA_CLEANUP_ENABLED=true \
     -e MEDIA_CLEANUP_RETENTION_DAYS=30 \
     whatsapp-local
   ```

3. **Verificar logs:**
   ```bash
   docker logs -f whatsapp | grep cleanup
   ```

---

**Versão:** 7.9.1 (com limpeza automática de mídia)
**Data:** 23 de Novembro de 2025
