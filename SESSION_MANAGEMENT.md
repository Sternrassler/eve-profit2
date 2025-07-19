# Session Management & Wiederaufnahme-Guide - EVE Profit Calculator 2.0

## ⚠️ **WICHTIG: Session-Limitationen verstehen**

### **Was passiert bei Unterbrechungen:**
GitHub Copilot kann **NICHT automatisch** Arbeit nach Session-Unterbrechungen wiederaufnehmen:

❌ **Session-Verlust bei:**
- Rechner geht aus / Neustart
- VS Code wird geschlossen
- Netzwerk-Unterbrechung
- Längere Inaktivität

❌ **Was verloren geht:**
- Mein "Gedächtnis" der aktuellen Arbeitsschritte
- Kontext der laufenden Implementation
- Tool-Call Historie
- Zwischenergebnisse

✅ **Was bleibt erhalten:**
- **Alle bereits erstellten/geänderten Dateien**
- **Komplette Projektdokumentation**
- **Implementierter Code**
- **Git-Historie (falls vorhanden)**

## 🔄 **Wiederaufnahme-Strategien**

### **1. Schnelle Wiederaufnahme (Standard-Verfahren):**
```
1. Neue Copilot-Session starten
2. Sagen: "Lies PROJECT_CONTEXT.md und DEVELOPMENT_GUIDELINES.md"
3. Kurz erklären: "Wir waren gerade dabei [X] zu implementieren"
4. Copilot analysiert vorhandene Dateien
5. Nahtlose Fortsetzung möglich
```

### **2. Status-basierte Wiederaufnahme:**
```
1. Vor Unterbrechung: "Erstelle einen Development-Checkpoint"
2. Nach Wiederaufnahme: "Lies STATUS.md und setze die Arbeit fort"
3. Copilot kann exakt da weitermachen wo aufgehört wurde
```

### **3. Datei-basierte Wiederaufnahme:**
```
1. Copilot analysiert zuletzt geänderte Dateien
2. Erkennt unvollständige Implementierungen
3. Fragt nach: "Soll ich [spezifische Aufgabe] fortsetzen?"
```

## 📋 **Checkpoint-System**

### **Wann Checkpoints erstellen:**
- **Vor längeren Arbeitsphasen** (>30 Minuten)
- **Nach Abschluss größerer Features**
- **Vor komplexen Refactoring-Arbeiten**
- **Am Ende des Arbeitstages**

### **Checkpoint-Befehl:**
```
"Erstelle einen Development-Checkpoint mit aktuellem Status"
```

### **Wiederaufnahme-Befehl:**
```
"Lies alle Projekt-Dokumente und STATUS.md, dann setze die Entwicklung fort"
```

## 🎯 **Best Practices für längere Entwicklungsphasen**

### **Arbeitsblöcke strukturieren:**
1. **30-45 Minuten Blöcke:** Überschaubare Implementierungs-Einheiten
2. **Zwischenspeichern:** Regelmäßige Status-Updates
3. **Dokumentieren:** Jeden Schritt in README oder STATUS festhalten
4. **Testen:** Funktionalität nach jedem Block prüfen

### **Empfohlene Arbeitssequenz:**
```
1. 🎯 Ziel definieren: "Implementiere ESI Client mit Character API"
2. 📝 Plan erstellen: "Welche Dateien/Funktionen sind nötig?"
3. ⏱️ Zeitschätzung: "Etwa 45 Minuten für Grundimplementierung"
4. 🔨 Implementation: Schritt-für-Schritt Umsetzung
5. 📊 Checkpoint: Status dokumentieren
6. ✅ Test: Funktionalität validieren
```

### **Signalwörter für Checkpoints:**
- "Mach eine Pause" → Checkpoint erstellen
- "Speichere den Fortschritt" → Status dokumentieren  
- "Das wars für heute" → Vollständiger Checkpoint
- "Morgen weiter" → Detaillierte Wiederaufnahme-Anleitung

## 📝 **Status-Dokumentation Template**

### **Git-Workflow für Entwicklungsphasen:**
```bash
# Nach jeder abgeschlossenen Phase automatischer Commit + Push:
./commit-phase.sh <phase-number> "<phase-name>" "[description]"

# Beispiele:
./commit-phase.sh 2 "SDE Client Implementation" "SQLite integration für Items, Stations, Regionen"
./commit-phase.sh 3 "ESI API Client" "Rate limiting und parallele Marktdaten-Abfragen"
./commit-phase.sh 4 "Character API" "EVE SSO OAuth + JWT Token Management"
```

### **Git-Historie als Fortschritts-Tracking:**
- Jeder Commit = Abgeschlossene Entwicklungsphase
- Detaillierte Commit-Messages mit Features + Tests
- Remote Repository: https://github.com/Sternrassler/eve-profit2
- Automatische Dokumentation des Entwicklungsfortschritts

### **Session-Wiederaufnahme via Git:**
```bash
git log --oneline              # Zeigt alle abgeschlossenen Phasen
git show HEAD                  # Details der letzten implementierten Phase
git diff HEAD~1                # Änderungen der letzten Phase
```

### **Automatisches Status-Update Format:**
```markdown
# Development Status - [Datum]

## ✅ Completed Today:
- [Liste der abgeschlossenen Features/Dateien]
- [Implementierte Funktionen]
- [Behobene Issues]

## 🚧 Currently Working On:
- [Aktuelle Aufgabe mit % Fortschritt]
- [Teilweise implementierte Features]
- [Offene Baustellen]

## 📁 Modified Files:
- `path/to/file.go` - [Was wurde geändert]
- `path/to/another.js` - [Implementierungsstand]

## ⏭️ Next Steps:
- [ ] [Nächste konkrete Aufgabe]
- [ ] [Folgeschritte]
- [ ] [Testing/Validation]

## 🐛 Known Issues:
- [Bekannte Probleme die noch zu lösen sind]

## 💡 Notes & Context:
- [Wichtige Entscheidungen]
- [Erkenntnisse während der Implementation]
- [Besondere Herausforderungen]

## 🔄 Resume Command:
"Lies alle Projekt-Dokumente und setze die Arbeit an [spezifische Aufgabe] fort"
```

## 🚀 **Wiederaufnahme-Befehle (Copy & Paste Ready)**

### **Für neue Sessions:**
```
Lies bitte PROJECT_CONTEXT.md, DEVELOPMENT_GUIDELINES.md und analysiere den aktuellen Projekt-Stand. Ich möchte die Entwicklung fortsetzen.
```

### **Mit Status-Datei:**
```
Lies alle Projektdokumente und STATUS.md, dann setze die Entwicklung genau dort fort wo wir aufgehört haben.
```

### **Für spezifische Features:**
```
Analysiere den aktuellen Stand der [Feature-Name] Implementation und setze die Arbeit daran fort. Zeige mir was bereits implementiert ist und was noch fehlt.
```

### **Nach Fehlern/Problemen:**
```
Analysiere die letzten Änderungen, prüfe auf Fehler und setze die Implementation fort. Falls Probleme vorliegen, behebe sie zuerst.
```

### **Für Testing-Phase:**
```
Prüfe alle implementierten Features auf Funktionalität und erstelle Tests für noch nicht getestete Bereiche.
```

## 🛡️ **Präventive Maßnahmen**

### **Vor jeder längeren Arbeitsphase:**
1. ✅ Aktueller Stand dokumentiert?
2. ✅ Alle Dateien gespeichert?
3. ✅ Klares Ziel definiert?
4. ✅ Zeitrahmen realistisch?

### **Während der Arbeit:**
- **Alle 20-30 Minuten:** Kurze Zwischenspeicherung
- **Bei komplexen Änderungen:** Sofortige Dokumentation
- **Vor Experimenten:** Backup-Punkt setzen

### **Bei Unsicherheit:**
```
"Erstelle jetzt einen Checkpoint bevor wir weitermachen"
```

## 🎯 **Spezielle Wiederaufnahme-Situationen**

### **Nach Hardware-Problemen:**
```
1. "Analysiere alle Dateien auf Integrität"
2. "Prüfe ob die letzte Implementation vollständig ist"
3. "Zeige mir den aktuellen Projekt-Status"
4. "Setze die Arbeit fort sobald alles validiert ist"
```

### **Nach längerer Pause (Tage/Wochen):**
```
1. "Gib mir eine komplette Projekt-Übersicht"
2. "Was ist implementiert, was fehlt noch?"
3. "Zeige mir die nächsten logischen Schritte"
4. "Frische mein Gedächtnis zu den Design-Entscheidungen auf"
```

### **Nach Konflikt/Merge-Problemen:**
```
1. "Analysiere alle Merge-Konflikte"
2. "Zeige mir was kaputt gegangen ist"
3. "Repariere systematisch alle Probleme"
4. "Validiere dass alles wieder funktioniert"
```

---

## 🎯 **Zusammenfassung: Die 3 goldenen Regeln**

1. **📝 Dokumentiere kontinuierlich** - Jeder Schritt wird festgehalten
2. **⏰ Arbeite in überschaubaren Blöcken** - Max. 45 Minuten ohne Checkpoint
3. **🔄 Nutze klare Wiederaufnahme-Befehle** - Copy & Paste aus diesem Guide

**Mit diesen Strategien ist eine verlustfreie Wiederaufnahme auch nach längeren Unterbrechungen möglich! 🚀**
