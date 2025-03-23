# 🧰 stamp

`stamp` ist ein schlankes Template-Tool in Go zur Verarbeitung von `.st`-Dateien mit Daten aus `.env`- und `.yaml`-Quellen.

## 🔧 Features

- Platzhalterersetzung mit `{{ .VARIABLEN }}`-Syntax
- Unterstützung für `.env` **und** `.yaml`-Dateien
- `--strict`: Fehler bei nicht gesetzten Variablen
- `--dry-run`: Ausgabe ins Terminal statt Datei
- `--self-test`: prüft dein Setup vorab
- Einfach als statisches Binary nutzbar (kein Docker nötig)

## 🚀 CLI-Beispiele

```bash
# Einzelnes Template rendern
stamp --render --in=config.tpl.st --out=config.conf --env=.env --yaml=config.yaml

# Im Batch alle .st-Dateien rendern
stamp --render --in-dir=templates/ --out-dir=out/ --env=.env --yaml=config.yaml

# Vorschau-Modus
stamp --render --in-dir=templates --dry-run

# Strikter Modus (Fehler bei fehlenden Variablen)
stamp --render --in=config.tpl.st --out=config.conf --strict

# Setup testen
stamp --self-test --in-dir=templates --out-dir=out --env=.env --yaml=config.yaml
```

## 🛠️ Build

```bash
go mod tidy
go build -o stamp main.go
```

Oder verwende die [GitHub Actions](.github/workflows/build.yml).

## 📄 Lizenz

[Apache 2.0](LICENSE)
