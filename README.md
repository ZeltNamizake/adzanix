# Adzanix

![Termux](https://img.shields.io/badge/Termux-success.svg?style=for-the-badge&color=grey&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB3aWR0aD0iNTEyIiBoZWlnaHQ9IjUxMiIgdmlld0JveD0iMCAwIDUxMiA1MTIiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI%2BCjxyZWN0IHdpZHRoPSI1MTIiIGhlaWdodD0iNTEyIiBmaWxsPSJibGFjayIvPgo8cGF0aCBkPSJNMTU2LjUgMTQyLjVIMTEzLjVMMTk4LjUgMjU1LjVWMjU2LjVMMTEzLjUgMzY5LjVIMTU2LjVMMjQxLjUgMjU2LjVWMjU1LjVMMTU2LjUgMTQyLjVaIiBmaWxsPSJ3aGl0ZSIvPgo8cGF0aCBkPSJNMzk4LjUgMzQxLjVIMjcwLjVWMzY5LjVIMzk4LjVWMzQxLjVaIiBmaWxsPSJ3aGl0ZSIvPgo8L3N2Zz4K)
![Linux](https://img.shields.io/badge/linux-badge?style=for-the-badge&logo=linux&color=black)

Adzanix is a minimal, no-distraction prayer time reminder for Termux/Linux, written in Go with automatic adhan playback.

---

## Features

- Automatic adhan playback
- Custom prayer schedule via config.json
- Timezone support

---

## Requirements

- Go (for building)
- mpv (for audio playback)

---

## Installation

```bash
git clone https://github.com/zeltnamizake/adzanix
cd adzanix
```

---

## Usage

1. Edit `config.json`:
   ```
   {
    "timezone": "Asia/Jakarta",
    "adzan_normal": "/path/to/adzan.mp3",
    "adzan_subuh": "/path/to/adzan_subuh.mp3",
    "pray_times": {
      "fajr": "04:45",
      "dhuhr": "12:05",
      "asr": "15:20",
      "maghrib": "18:10",
      "isha": "19:20"
     }
   }
   ```

2. Run:
   ```bash
   go run main.go
   ```

---

## Notes

- Make sure `mpv` is installed:
  ```bash
  apt install mpv
  ```
- Run in the background:
  ```bash
  go run main.go &
  ```

