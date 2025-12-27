# 🔐 GoBruteSim – Brute Force Attack Simulator in Go

> **GoBruteSim** is an educational brute-force password attack **simulation** written in Go.
> It demonstrates **concurrency, worker pools, context cancellation, and channel-based communication** using a modular project structure.

⚠️ **Disclaimer:**
This project is for **learning and demonstration purposes only**. It is **not intended for real-world password cracking or malicious use**.

---

## 📂 Project Structure

```
.
├── attack/
│   ├── config.go
│   ├── doc.go
│   ├── engine.go
│   ├── result.go
│   ├── stats.go
│   └── worker.go
│
├── docs/
│   ├── attack.html
│   ├── attack.txt
│   ├── miniproject.html
│   ├── miniproject.txt
│   ├── utils.html
│   └── utils.txt
│
├── logs/
│   └── app.log
│
├── utils/
│   ├── doc.go
│   ├── impfunc.go
│   ├── logger.go
│   ├── mask.go
│   └── password.go
│
├── main.go
├── go.mod
└── run.sh

```
### 📌 Notes

* `attack/` – Core brute-force simulation engine
* `utils/` – Logging, hashing, mask & password helpers
* `docs/` – Generated documentation (HTML & TXT)
* `logs/` – Runtime log files
* `main.go` – Application entry point
* `run.sh` – Convenience script to run the project

---

## 🚀 Features

* Concurrent brute-force simulation using **goroutines**
* Configurable worker pool
* Mask-based password generation
* Hash comparison & encoding utilities
* Context-based cancellation
* Runtime statistics and attack results
* Custom logging levels (`info`, `debug`)

---

## 🧠 How It Works

1. **Main (`miniproject`)**

   * Initializes logging
   * Configures the attack parameters
   * Starts the brute-force engine

2. **Attack Package**

   * `AttackConfig` – defines charset, mask, workers, and target hash
   * `Engine` – manages workers and job distribution
   * `Worker` – attempts passwords and compares hashes
   * `AttackStats` & `AttackResult` – track runtime metrics

3. **Utils Package**

   * Logging helpers
   * Hash generation and comparison
   * Password generation (random & mask-based)
   * Worker ID tracking

---

## 🛠️ Usage

### Run with logging level

```bash
go run . info
```

```bash
go run . debug
```

* `info` → normal runtime output
* `debug` → detailed worker-level logs

---

## 📦 Packages Overview

### `attack`

Handles the brute-force simulation logic.

Key types:

* `AttackConfig`
* `Engine`
* `Worker`
* `AttackStats`
* `AttackResult`

Designed to showcase **concurrent worker patterns** in Go.

---

### `utils`

Utility helpers used across the project.

Includes:

* Logging (`Init`, `Infof`, `Debugf`, `Warnf`, `Errorlog`)
* Hashing (`ConvertToHash`, `CompareWithHash`)
* Password generators (`GenerateFromMask`, `GenerateRandomPassword`)
* Misc helpers (`GetWorkerid`, Base64 encoding)

---

## 🎓 Learning Objectives

This project is ideal for understanding:

* Goroutines and channels
* Worker pool patterns
* Context cancellation
* Modular Go project layout
* Writing `godoc`-friendly packages

---

## 📌 Future Improvements (Optional)

* Support for multiple hash algorithms
* CLI flags instead of positional arguments
* Performance benchmarking
* Result export (JSON/CSV)
* Rate limiting simulation

---

## 📝 License

This project is released for **educational use only**.
Use responsibly.

---
