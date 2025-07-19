# EVE ESI Integration Dokumentation

## ESI Application Settings

Das EVE Profit Calculator Backend ist für die Integration mit der EVE ESI (EVE Swagger Interface) konfiguriert. Hier sind die wichtigsten Konfigurationsdetails:

### Application Registration
- **Name:** EVE Profit Maximizer
- **Beschreibung:** Tool für realistische Arbitrage-Berechnungen
- **Client ID:** `0928b4bcd20242aeb9b8be10f5451094`
- **Client Secret:** `AQPjLZ3VYAewR59J5jStZs52dY7jISGVLwXv5NA`
- **Callback URL:** `http://localhost:9000/callback`

### Server Konfiguration
- **Backend Port:** 9000 (geändert von 8081)
- **Frontend Dev Port:** 3000
- **API Base URL:** `http://localhost:9000`

### ESI Scopes (Berechtigungen)
Das Backend ist für folgende ESI Scopes konfiguriert:

1. **publicData** - Grundlegende öffentliche Informationen
2. **esi-location.read_location.v1** - Aktuelle Standort-Informationen
3. **esi-location.read_ship_type.v1** - Informationen zum aktuellen Schiff
4. **esi-skills.read_skills.v1** - Charakter-Fähigkeiten
5. **esi-wallet.read_character_wallet.v1** - Wallet-Informationen
6. **esi-universe.read_structures.v1** - Struktur-Informationen
7. **esi-assets.read_assets.v1** - Asset-Informationen
8. **esi-fittings.read_fittings.v1** - Schiffs-Fittings
9. **esi-characters.read_standings.v1** - Standings zu anderen Entitäten

### Konfigurationsdateien

#### .env.example
Enthält alle Umgebungsvariablen mit Standardwerten:
```env
ESI_CLIENT_ID=0928b4bcd20242aeb9b8be10f5451094
ESI_CLIENT_SECRET=AQPjLZ3VYAewR59J5jStZs52dY7jISGVLwXv5NA
ESI_CALLBACK_URL=http://localhost:9000/callback
ESI_SCOPES=publicData esi-location.read_location.v1 ...
```

#### internal/config/config.go
Zentrales Konfigurationsmanagement mit Umgebungsvariablen-Support und sinnvollen Defaults.

### Rate Limiting
- **Basis Rate Limit:** 150 Anfragen/Sekunde (ESI Standard)
- **Burst Limit:** 400 Anfragen
- **Timeout:** 30 Sekunden
- **Retry Logic:** 3 Versuche bei Fehlern

### Caching Strategien
- **Market Orders:** 5 Minuten TTL
- **Market History:** 1 Stunde TTL  
- **Type Info:** 24 Stunden TTL
- **Character Info:** 30 Minuten TTL

### Wichtige URLs
- **ESI Base:** `https://esi.evetech.net`
- **EVE SSO:** `https://login.eveonline.com`
- **Authorization:** `https://login.eveonline.com/v2/oauth/authorize`
- **Token:** `https://login.eveonline.com/v2/oauth/token`

### Development Setup
1. Kopiere `.env.example` zu `.env`
2. Passe bei Bedarf die Konfiguration an
3. Starte den Server mit `go run cmd/server/main.go`
4. Server läuft auf Port 9000 (nicht mehr 8081)

### Sicherheitshinweise
- **Client Secret** niemals in Git committen
- Für Produktion separate EVE Application registrieren
- HTTPS für Production verwenden
- Callback URL entsprechend anpassen

### API Endpoints (geplant)
- `GET /api/v1/auth/login` - EVE SSO Login initiieren
- `GET /callback` - OAuth Callback Handler
- `POST /api/v1/auth/refresh` - Token Refresh
- `GET /api/v1/character/info` - Character-Informationen

Die ESI Integration ermöglicht es, authentifizierte Anfragen an die EVE API zu stellen und Character-spezifische Daten abzurufen.
