# 🕵️‍♂️ Go Multi-Port- und Netzwerk-Scanner

Ein schneller, schlanker TCP-Portscanner in Go, der:

- 🔍 Alle lokalen IPv4-Adressen erkennt (inkl. `127.0.0.1`, `192.168.x.x`, `10.x.x.x` etc.)
- 🌐 Das gesamte lokale Netzwerk (z. B. `192.168.178.0/24`) scannt
- ⚡ Mit Goroutinen arbeitet (Multithreading)
- 🔒 Häufig genutzte Ports auf jedem Host testet

---

## 📦 Features

- Automatische Subnetz-Erkennung
- Scan beliebter Ports wie `22`, `80`, `443`, `3306`, `8080`, ...
- Gleichzeitiges Scannen von bis zu 254 Hosts (Multithreaded)
- Optimiert für Geschwindigkeit durch `net.DialTimeout`

---

## 🧪 Beispielausgabe

```
Gefundenes Subnetz: 192.168.178.0/24
[+] 192.168.178.1: Offen -> [80]
[+] 192.168.178.22: Offen -> [22 443]
[+] 192.168.178.101: Offen -> [3306]
Scan abgeschlossen.
```

---

## ⚙️ Installation

### Voraussetzung

- [Go](https://go.dev/dl/) (Version ≥ 1.16 empfohlen)

### Kompilieren

```bash
git clone https://github.com/deinuser/multiscanner.git
cd multiscanner
go build -o multiscanner multiscanner.go
```

---

## 🚀 Verwendung

```bash
./multiscanner
```

Scant dein gesamtes Subnetz (`/24`) auf gängige Ports. Keine Root-Rechte nötig.

---

## 🔧 Anpassung

Wenn du mehr Kontrolle willst, kannst du Folgendes direkt im Code anpassen:

### 1. Portliste erweitern

```go
var commonPorts = []int{22, 80, 443, 3306, 8080, 8443, 5900, 21, 25}
```

### 2. Timeout ändern

```go
var timeout = 300 * time.Millisecond
```

---

## 📁 Projektstruktur

```text
multiscanner/
├── multiscanner.go   # Hauptprogramm
└── README.md         # Dokumentation
```

---

## 🛡️ Sicherheitshinweis

Dieser Scanner ist für **lokale** Netzwerke gedacht. Der Einsatz in fremden Netzen ohne Erlaubnis kann strafbar sein (z. B. §202c StGB in Deutschland).

---

## 🧑‍💻 Autor

**Marcus Dziersan**  
Projektname: **TwistedScan**  
Lizenz: MIT

---

## 📜 Lizenz

MIT License – freie Nutzung, aber ohne Garantie.
