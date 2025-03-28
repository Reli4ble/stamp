# 🧰 stamp

`stamp` ist ein schlankes Template-Tool in Go zur Verarbeitung von `.st`-Dateien mit Daten aus `.env`- und `.yaml`-Quellen. Es unterstützt den strikten Modus (Fehler bei fehlenden Platzhaltern), Probelauf (Ausgabevorschau im Terminal), Selbsttest, Stapelverarbeitung und mehr.

## 🔧 Features

- **Platzhalterersetzung:** Verwenden Sie die Syntax `{{ .VARIABLE }}` in Ihren Vorlagen.
- **Datenquellen:** Führen Sie Daten aus einer `.env`-Datei und einer YAML-Datei zusammen (YAML-Werte überschreiben ENV im Konfliktfall).
- **Strikter Modus:** Fehlerausgabe bei fehlendem Platzhalter.
- **Probelaufmodus:** Zeigen Sie die gerenderte Vorlage im Terminal an, anstatt sie in eine Datei zu schreiben.
- **Stapelverarbeitung:** Verarbeiten Sie alle `.st`-Dateien in einem Verzeichnis und geben Sie sie aus (ohne die Erweiterung `.st`).
- **Plattformübergreifende Builds:** Erstellt über GitHub Actions für Linux und macOS (amd64 & arm64).
- **Selbsttestmodus:** Überprüfen Sie Ihre Konfiguration, prüfen Sie, ob erforderliche Dateien vorhanden sind, und stellen Sie die Schreibberechtigungen sicher.
- **Automatischer Scanmodus:** Rendern Sie alle Dateien in einem Verzeichnis rekursiv und überschreiben Sie sie an Ort und Stelle.
- **Beenden-Verhalten steuern:** mit --force-success.

# 🧰 stamp

stamp is a lightweight template tool written in Go that processes `.st` files using variables provided in `.env` and `.yaml` files. It supports strict mode (errors on missing placeholders), dry-run (output preview in the terminal), self-test, batch processing and more.

## 🔧 Features

- **Placeholder Replacement:** Use `{{ .VARIABLE }}` syntax in your templates.
- **Data Sources:** Merge data from a `.env` file and a YAML file (YAML values override ENV in case of conflicts).
- **Strict Mode:** Error out if a placeholder is missing.
- **Dry-run Mode:** Display the rendered template in the terminal instead of writing to a file.
- **Batch Processing:** Process all `.st` files in a directory and output them (without the `.st` extension).
- **Cross-Platform Builds:** Built via GitHub Actions for Linux, and macOS (amd64 & arm64).
- **Self-Test Mode:** Validate your configuration, check for required files, and ensure write permissions.
- **Auto-scan mode:** Recursively render all files in a directory and overwrite them in place.
- **Control exit behavior:** with --force-success.




## 📄 Lizenz

[Apache 2.0](LICENSE)
