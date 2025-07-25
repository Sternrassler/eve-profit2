# Project Security Guidelines

## 🔒 Sichere Konfiguration für EVE ESI Integration

### Warum diese Konfiguration wichtig ist
Der EVE ESI Client Secret ist ein **kritischer Sicherheitsschlüssel**, der niemals in öffentlichen Repositories gespeichert werden darf. Ein kompromittierter Client Secret kann zu unauthorisiertem Zugriff auf EVE Online Spielerdaten führen.

## 🚀 Setup für Entwicklung

### 1. Lokale .env Datei erstellen
```bash
# Kopiere die Vorlage
cp backend/.env.example backend/.env

# Bearbeite die .env Datei mit echten Werten
nano backend/.env
```

**💡 Vollständige API Security Configuration siehe: `PROJECT_API_SPECS.md`**

## 🏭 Production Deployment

**💡 Vollständige Production Deployment Commands siehe: `PROJECT_API_SPECS.md`**

### Azure Key Vault (Enterprise)
```go
// Beispiel Go-Code für Azure Key Vault
import "github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"

func getESISecret(ctx context.Context) (string, error) {
    client := azsecrets.NewClient("https://myvault.vault.azure.net/", cred, nil)
    resp, err := client.GetSecret(ctx, "esi-client-secret", "", nil)
    if err != nil {
        return "", err
    }
    return *resp.Value, nil
}
```

## 🛡️ Sicherheits-Checklist

### ✅ MUST DO:
- [ ] **.env in .gitignore** (bereits erledigt)
- [ ] **.env.example nur mit Dummy-Werten** (bereits erledigt)
- [ ] **Echte Credentials nur in lokaler .env oder Environment Variables**
- [ ] **Production: Key Vault oder Secret Management Service**

### ❌ NEVER DO:
- ❌ Client Secret in .env.example committen
- ❌ .env Dateien in Git committen
- ❌ Screenshots mit Credentials teilen
- ❌ Client Secret in Code hardcoden
- ❌ Credentials in Logs ausgeben

## 🔧 Development Team Setup

### Für neue Team-Mitglieder:
```bash
# 1. Repository klonen
git clone https://github.com/Sternrassler/eve-profit2.git

# 2. Environment setup
cd eve-profit2/backend
cp .env.example .env

# 3. Credentials vom Team Lead erhalten
# - ESI_CLIENT_ID und ESI_CLIENT_SECRET
# - Oder eigene EVE Application erstellen

# 4. Server starten
go run cmd/server/main.go
```

### Team-Credentials Sharing (Sicher):
1. **Passwort-Manager** (1Password, Bitwarden) für Team
2. **Verschlüsselte Chat** (Signal, secure Slack)
3. **Persönliche EVE Applications** für jeden Entwickler
4. **Separate Applications** für Development/Staging/Production

## 🚨 Was tun bei Credential-Leak?

### Sofortige Maßnahmen:
1. **EVE Application deaktivieren** in EVE Developer Portal
2. **Neue Application erstellen** mit neuen Credentials
3. **Git History bereinigen** (falls committed)
4. **Team benachrichtigen** über neue Credentials

### Git History bereinigen:
```bash
# Gefährlich! Backup zuerst!
git filter-branch --force --index-filter \
  'git rm --cached --ignore-unmatch backend/.env' \
  --prune-empty --tag-name-filter cat -- --all

# Oder mit BFG Repo-Cleaner (sicherer)
java -jar bfg.jar --delete-files .env
git reflog expire --expire=now --all && git gc --prune=now --aggressive
```

---

## 🎯 Für dieses Projekt

**Aktuelle EVE Application:**
- **Client ID:** `your_client_id_here` (public OK)
- **Client Secret:** `your_client_secret_here` (PRIVATE!)
- **Status:** Credentials aus .env.example entfernt ✅

**Nächste Schritte:**
1. Lokale .env mit echten Werten erstellen
2. Server testen: `go run cmd/server/main.go`
3. ESI Test: `curl http://localhost:9000/api/v1/esi/test`
