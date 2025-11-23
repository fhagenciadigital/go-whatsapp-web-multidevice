# Script para construir e publicar a imagem Docker no Docker Hub
# Usage: .\build-and-push.ps1 -DockerHubUsername "seu-usuario"

param(
    [Parameter(Mandatory=$true)]
    [string]$DockerHubUsername,
    
    [Parameter(Mandatory=$false)]
    [string]$ImageName = "go-whatsapp-web-multidevice",
    
    [Parameter(Mandatory=$false)]
    [string]$Tag = "latest"
)

$ErrorActionPreference = "Stop"

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Build and Push WhatsApp Docker Image" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# Verificar se está no diretório correto
$ProjectRoot = "c:\Users\paulo\repos\go-whatsapp-web-multidevice"
if (-not (Test-Path "$ProjectRoot\docker\golang.Dockerfile")) {
    Write-Host "ERROR: Dockerfile não encontrado!" -ForegroundColor Red
    exit 1
}

Write-Host "📦 Projeto: $ProjectRoot" -ForegroundColor Green
Write-Host "🐳 Docker Hub: $DockerHubUsername/$ImageName`:$Tag" -ForegroundColor Green
Write-Host ""

# 1. Construir a imagem
Write-Host "🔨 Passo 1: Construindo a imagem Docker..." -ForegroundColor Yellow
Write-Host ""

$BuildCommand = "docker build -f docker/golang.Dockerfile -t ${DockerHubUsername}/${ImageName}:${Tag} ."
Write-Host "Executando: $BuildCommand" -ForegroundColor Gray

Set-Location $ProjectRoot
docker build -f docker/golang.Dockerfile -t "${DockerHubUsername}/${ImageName}:${Tag}" .

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Erro ao construir a imagem!" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "✅ Imagem construída com sucesso!" -ForegroundColor Green
Write-Host ""

# 2. Verificar se está autenticado no Docker Hub
Write-Host "🔐 Passo 2: Verificando autenticação no Docker Hub..." -ForegroundColor Yellow

$DockerInfo = docker info 2>&1 | Out-String
if ($DockerInfo -notmatch $DockerHubUsername) {
    Write-Host "⚠️  Você não está autenticado no Docker Hub" -ForegroundColor Yellow
    Write-Host "Por favor, faça login:" -ForegroundColor Yellow
    Write-Host ""
    docker login
    
    if ($LASTEXITCODE -ne 0) {
        Write-Host "❌ Erro ao fazer login no Docker Hub!" -ForegroundColor Red
        exit 1
    }
}

Write-Host "✅ Autenticado no Docker Hub" -ForegroundColor Green
Write-Host ""

# 3. Enviar para o Docker Hub
Write-Host "📤 Passo 3: Enviando imagem para o Docker Hub..." -ForegroundColor Yellow
Write-Host ""

docker push "${DockerHubUsername}/${ImageName}:${Tag}"

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Erro ao enviar a imagem!" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Green
Write-Host "✅ SUCESSO!" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Green
Write-Host ""
Write-Host "Sua imagem foi publicada com sucesso!" -ForegroundColor Green
Write-Host ""
Write-Host "📦 Imagem: ${DockerHubUsername}/${ImageName}:${Tag}" -ForegroundColor Cyan
Write-Host ""
Write-Host "Para usar a imagem:" -ForegroundColor Yellow
Write-Host "  docker pull ${DockerHubUsername}/${ImageName}:${Tag}" -ForegroundColor White
Write-Host "  docker run -p 3000:3000 ${DockerHubUsername}/${ImageName}:${Tag}" -ForegroundColor White
Write-Host ""
