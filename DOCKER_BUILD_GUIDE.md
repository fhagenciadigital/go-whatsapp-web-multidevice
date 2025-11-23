# 🐳 Guia: Construir e Publicar Imagem Docker

Este guia explica como construir a imagem Docker com as suas alterações e publicá-la no Docker Hub.

## 📋 Pré-requisitos

1. Docker instalado e em execução
2. Conta no Docker Hub (https://hub.docker.com)
3. Saber o seu nome de usuário do Docker Hub

## 🚀 Método 1: Usando o Script PowerShell (Recomendado)

### Passo 1: Execute o script

```powershell
cd c:\Users\paulo\repos\go-whatsapp-web-multidevice

# Substitua "seu-usuario" pelo seu nome de usuário do Docker Hub
.\build-and-push.ps1 -DockerHubUsername "seu-usuario"
```

### Opções adicionais:

```powershell
# Com nome de imagem e tag personalizados
.\build-and-push.ps1 -DockerHubUsername "seu-usuario" -ImageName "whatsapp-api" -Tag "v1.0"
```

O script irá:
- ✅ Construir a imagem Docker
- ✅ Fazer login no Docker Hub (se necessário)
- ✅ Enviar a imagem para o Docker Hub
- ✅ Mostrar comandos para usar a imagem

---

## 🛠️ Método 2: Comandos Manuais

### Passo 1: Fazer login no Docker Hub

```powershell
docker login
```

Insira seu nome de usuário e senha quando solicitado.

### Passo 2: Construir a imagem

```powershell
cd c:\Users\paulo\repos\go-whatsapp-web-multidevice

# Substitua "seu-usuario" pelo seu nome de usuário do Docker Hub
docker build -f docker/golang.Dockerfile -t seu-usuario/go-whatsapp-web-multidevice:latest .
```

### Passo 3: Enviar para o Docker Hub

```powershell
docker push seu-usuario/go-whatsapp-web-multidevice:latest
```

---

## 📦 Como Usar a Imagem Publicada

### Opção 1: Docker Run

```bash
docker run -d \
  --name whatsapp \
  -p 3000:3000 \
  -v whatsapp-data:/app/storages \
  seu-usuario/go-whatsapp-web-multidevice:latest
```

### Opção 2: Docker Compose

1. Edite o arquivo `docker-compose.custom.yml`
2. Substitua `YOUR_DOCKERHUB_USERNAME` pelo seu nome de usuário
3. Execute:

```powershell
docker-compose -f docker-compose.custom.yml up -d
```

---

## 🔧 Testando a Imagem Localmente (Antes de Publicar)

Se quiser testar a imagem localmente antes de enviar para o Docker Hub:

```powershell
# Construir apenas
docker build -f docker/golang.Dockerfile -t whatsapp-test:local .

# Executar
docker run -d -p 3000:3000 --name whatsapp-test whatsapp-test:local

# Testar
# Abra: http://localhost:3000

# Limpar
docker stop whatsapp-test
docker rm whatsapp-test
docker rmi whatsapp-test:local
```

---

## 📝 Versionamento de Imagens

É uma boa prática usar tags de versão:

```powershell
# Construir com múltiplas tags
docker build -f docker/golang.Dockerfile \
  -t seu-usuario/go-whatsapp-web-multidevice:latest \
  -t seu-usuario/go-whatsapp-web-multidevice:v1.0.0 \
  -t seu-usuario/go-whatsapp-web-multidevice:audio-fix \
  .

# Enviar todas as tags
docker push seu-usuario/go-whatsapp-web-multidevice:latest
docker push seu-usuario/go-whatsapp-web-multidevice:v1.0.0
docker push seu-usuario/go-whatsapp-web-multidevice:audio-fix
```

---

## 🔍 Verificar Imagens

```powershell
# Listar imagens locais
docker images | Select-String "whatsapp"

# Ver detalhes da imagem
docker inspect seu-usuario/go-whatsapp-web-multidevice:latest

# Ver tamanho da imagem
docker images seu-usuario/go-whatsapp-web-multidevice --format "{{.Size}}"
```

---

## 🆘 Solução de Problemas

### Erro: "denied: requested access to the resource is denied"

**Solução:** Você não está autenticado ou não tem permissão. Execute:

```powershell
docker logout
docker login
```

### Erro: "Cannot connect to the Docker daemon"

**Solução:** Certifique-se de que o Docker Desktop está em execução.

### Erro durante o build

**Solução:** Verifique se todos os arquivos estão no lugar:

```powershell
Test-Path "c:\Users\paulo\repos\go-whatsapp-web-multidevice\docker\golang.Dockerfile"
Test-Path "c:\Users\paulo\repos\go-whatsapp-web-multidevice\src\go.mod"
```

### Imagem muito grande

**Solução:** O Dockerfile já usa multi-stage build para otimizar o tamanho. A imagem final deve ter aproximadamente 80-100 MB.

---

## 📊 Exemplo Completo

```powershell
# 1. Navegar para o projeto
cd c:\Users\paulo\repos\go-whatsapp-web-multidevice

# 2. Fazer login
docker login

# 3. Construir a imagem
docker build -f docker/golang.Dockerfile -t paulouser/whatsapp-api:audio-fix .

# 4. Testar localmente
docker run -d -p 3000:3000 --name test-whatsapp paulouser/whatsapp-api:audio-fix

# 5. Verificar se funciona
Start-Process "http://localhost:3000"

# 6. Se estiver OK, enviar para Docker Hub
docker push paulouser/whatsapp-api:audio-fix

# 7. Limpar teste
docker stop test-whatsapp
docker rm test-whatsapp
```

---

## ✅ Checklist Final

- [ ] Docker Desktop está em execução
- [ ] Fez login no Docker Hub (`docker login`)
- [ ] Construiu a imagem com sucesso
- [ ] Testou a imagem localmente (opcional mas recomendado)
- [ ] Enviou a imagem para o Docker Hub
- [ ] Verificou que a imagem aparece no Docker Hub (https://hub.docker.com)

---

## 🔗 Links Úteis

- Docker Hub: https://hub.docker.com
- Docker Documentation: https://docs.docker.com
- Projeto Original: https://github.com/aldinokemal/go-whatsapp-web-multidevice
