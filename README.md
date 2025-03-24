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

## 📄 Lizenz

[Apache 2.0](LICENSE)


# 🧰 stamp

stamp is a lightweight template tool written in Go that processes `.st` files using variables provided in `.env` and `.yaml` files. It supports strict mode (errors on missing placeholders), dry-run (output preview in the terminal), self-test, and batch processing.

## 🔧 Features

- **Placeholder Replacement:** Use `{{ .VARIABLE }}` syntax in your templates.
- **Data Sources:** Merge data from a `.env` file and a YAML file (YAML values override ENV in case of conflicts).
- **Strict Mode:** Error out if a placeholder is missing.
- **Dry-run Mode:** Display the rendered template in the terminal instead of writing to a file.
- **Batch Processing:** Process all `.st` files in a directory and output them (without the `.st` extension).
- **Cross-Platform Builds:** Built via GitHub Actions for Linux, and macOS (amd64 & arm64).
- **Self-Test Mode:** Validate your configuration, check for required files, and ensure write permissions.

## Installation

Clone the repository and build stamp:

```bash
git clone https://github.com/Reliable/stamp.git
cd stamp
go mod tidy
go build -o stamp main.go


## 📄 Lizenz

[Apache 2.0](LICENSE)
