# Note rest api service

---

[![forthebadge](data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNjEuMzYiIGhlaWdodD0iMzUiIHZpZXdCb3g9IjAgMCAxNjEuMzYgMzUiPjxyZWN0IGNsYXNzPSJzdmdfX3JlY3QiIHg9IjAiIHk9IjAiIHdpZHRoPSIxMTUuMzEiIGhlaWdodD0iMzUiIGZpbGw9IiMzMUM0RjMiLz48cmVjdCBjbGFzcz0ic3ZnX19yZWN0IiB4PSIxMTMuMzEiIHk9IjAiIHdpZHRoPSI0OC4wNTAwMDAwMDAwMDAwMSIgaGVpZ2h0PSIzNSIgZmlsbD0iIzM4OUFENSIvPjxwYXRoIGNsYXNzPSJzdmdfX3RleHQiIGQ9Ik0xNS42OSAyMkwxNC4yMiAyMkwxNC4yMiAxMy40N0wxNi4xNCAxMy40N0wxOC42MCAyMC4wMUwyMS4wNiAxMy40N0wyMi45NyAxMy40N0wyMi45NyAyMkwyMS40OSAyMkwyMS40OSAxOS4xOUwyMS42NCAxNS40M0wxOS4xMiAyMkwxOC4wNiAyMkwxNS41NSAxNS40M0wxNS42OSAxOS4xOUwxNS42OSAyMlpNMjguNDkgMjJMMjYuOTUgMjJMMzAuMTcgMTMuNDdMMzEuNTAgMTMuNDdMMzQuNzMgMjJMMzMuMTggMjJMMzIuNDkgMjAuMDFMMjkuMTggMjAuMDFMMjguNDkgMjJaTTMwLjgzIDE1LjI4TDI5LjYwIDE4LjgyTDMyLjA3IDE4LjgyTDMwLjgzIDE1LjI4Wk00MS4xNCAyMkwzOC42OSAyMkwzOC42OSAxMy40N0w0MS4yMSAxMy40N1E0Mi4zNCAxMy40NyA0My4yMSAxMy45N1E0NC4wOSAxNC40OCA0NC41NyAxNS40MFE0NS4wNSAxNi4zMyA0NS4wNSAxNy41Mkw0NS4wNSAxNy41Mkw0NS4wNSAxNy45NVE0NS4wNSAxOS4xNiA0NC41NyAyMC4wOFE0NC4wOCAyMS4wMCA0My4xOSAyMS41MFE0Mi4zMCAyMiA0MS4xNCAyMkw0MS4xNCAyMlpNNDAuMTcgMTQuNjZMNDAuMTcgMjAuODJMNDEuMTQgMjAuODJRNDIuMzAgMjAuODIgNDIuOTMgMjAuMDlRNDMuNTUgMTkuMzYgNDMuNTYgMTcuOTlMNDMuNTYgMTcuOTlMNDMuNTYgMTcuNTJRNDMuNTYgMTYuMTMgNDIuOTYgMTUuNDBRNDIuMzUgMTQuNjYgNDEuMjEgMTQuNjZMNDEuMjEgMTQuNjZMNDAuMTcgMTQuNjZaTTU1LjA5IDIyTDQ5LjUxIDIyTDQ5LjUxIDEzLjQ3TDU1LjA1IDEzLjQ3TDU1LjA1IDE0LjY2TDUxLjAwIDE0LjY2TDUxLjAwIDE3LjAyTDU0LjUwIDE3LjAyTDU0LjUwIDE4LjE5TDUxLjAwIDE4LjE5TDUxLjAwIDIwLjgyTDU1LjA5IDIwLjgyTDU1LjA5IDIyWk02Ni42NSAyMkw2NC42OCAxMy40N0w2Ni4xNSAxMy40N0w2Ny40NyAxOS44OEw2OS4xMCAxMy40N0w3MC4zNCAxMy40N0w3MS45NiAxOS44OUw3My4yNyAxMy40N0w3NC43NCAxMy40N0w3Mi43NyAyMkw3MS4zNSAyMkw2OS43MyAxNS43N0w2OC4wNyAyMkw2Ni42NSAyMlpNODAuMzggMjJMNzguOTAgMjJMNzguOTAgMTMuNDdMODAuMzggMTMuNDdMODAuMzggMjJaTTg2Ljg3IDE0LjY2TDg0LjIzIDE0LjY2TDg0LjIzIDEzLjQ3TDkxLjAwIDEzLjQ3TDkxLjAwIDE0LjY2TDg4LjM0IDE0LjY2TDg4LjM0IDIyTDg2Ljg3IDIyTDg2Ljg3IDE0LjY2Wk05Ni4yNCAyMkw5NC43NSAyMkw5NC43NSAxMy40N0w5Ni4yNCAxMy40N0w5Ni4yNCAxNy4wMkwxMDAuMDUgMTcuMDJMMTAwLjA1IDEzLjQ3TDEwMS41MyAxMy40N0wxMDEuNTMgMjJMMTAwLjA1IDIyTDEwMC4wNSAxOC4yMUw5Ni4yNCAxOC4yMUw5Ni4yNCAyMloiIGZpbGw9IiNGRkZGRkYiLz48cGF0aCBjbGFzcz0ic3ZnX190ZXh0IiBkPSJNMTI3LjA3IDE3LjgwTDEyNy4wNyAxNy44MFExMjcuMDcgMTYuNTQgMTI3LjY3IDE1LjU0UTEyOC4yNyAxNC41NSAxMjkuMzMgMTMuOTlRMTMwLjQwIDEzLjQzIDEzMS43NSAxMy40M0wxMzEuNzUgMTMuNDNRMTMyLjkyIDEzLjQzIDEzMy44NiAxMy44M1ExMzQuODAgMTQuMjIgMTM1LjQyIDE0Ljk3TDEzNS40MiAxNC45N0wxMzMuOTEgMTYuMzNRMTMzLjA2IDE1LjQwIDEzMS44OSAxNS40MEwxMzEuODkgMTUuNDBRMTMxLjg3IDE1LjQwIDEzMS44NyAxNS40MEwxMzEuODcgMTUuNDBRMTMwLjc5IDE1LjQwIDEzMC4xMyAxNi4wNlExMjkuNDcgMTYuNzEgMTI5LjQ3IDE3LjgwTDEyOS40NyAxNy44MFExMjkuNDcgMTguNTAgMTI5Ljc3IDE5LjA0UTEzMC4wNyAxOS41OSAxMzAuNjEgMTkuODlRMTMxLjE1IDIwLjIwIDEzMS44NSAyMC4yMEwxMzEuODUgMjAuMjBRMTMyLjUzIDIwLjIwIDEzMy4xMyAxOS45M0wxMzMuMTMgMTkuOTNMMTMzLjEzIDE3LjYyTDEzNS4yMyAxNy42MkwxMzUuMjMgMjEuMTBRMTM0LjUxIDIxLjYxIDEzMy41NyAyMS44OVExMzIuNjQgMjIuMTcgMTMxLjcwIDIyLjE3TDEzMS43MCAyMi4xN1ExMzAuMzggMjIuMTcgMTI5LjMyIDIxLjYxUTEyOC4yNyAyMS4wNSAxMjcuNjcgMjAuMDVRMTI3LjA3IDE5LjA2IDEyNy4wNyAxNy44MFpNMTM5Ljc5IDE3LjgwTDEzOS43OSAxNy44MFExMzkuNzkgMTYuNTUgMTQwLjM5IDE1LjU1UTE0MS4wMCAxNC41NiAxNDIuMDYgMTQuMDBRMTQzLjEyIDEzLjQzIDE0NC40NSAxMy40M0wxNDQuNDUgMTMuNDNRMTQ1Ljc4IDEzLjQzIDE0Ni44NSAxNC4wMFExNDcuOTEgMTQuNTYgMTQ4LjUyIDE1LjU1UTE0OS4xMiAxNi41NSAxNDkuMTIgMTcuODBMMTQ5LjEyIDE3LjgwUTE0OS4xMiAxOS4wNSAxNDguNTIgMjAuMDRRMTQ3LjkxIDIxLjA0IDE0Ni44NSAyMS42MFExNDUuNzkgMjIuMTcgMTQ0LjQ1IDIyLjE3TDE0NC40NSAyMi4xN1ExNDMuMTIgMjIuMTcgMTQyLjA2IDIxLjYwUTE0MS4wMCAyMS4wNCAxNDAuMzkgMjAuMDRRMTM5Ljc5IDE5LjA1IDEzOS43OSAxNy44MFpNMTQyLjE5IDE3LjgwTDE0Mi4xOSAxNy44MFExNDIuMTkgMTguNTEgMTQyLjQ5IDE5LjA1UTE0Mi43OSAxOS42MCAxNDMuMzEgMTkuOTBRMTQzLjgyIDIwLjIwIDE0NC40NSAyMC4yMEwxNDQuNDUgMjAuMjBRMTQ1LjA5IDIwLjIwIDE0NS42MSAxOS45MFExNDYuMTIgMTkuNjAgMTQ2LjQyIDE5LjA1UTE0Ni43MiAxOC41MSAxNDYuNzIgMTcuODBMMTQ2LjcyIDE3LjgwUTE0Ni43MiAxNy4wOSAxNDYuNDIgMTYuNTRRMTQ2LjEyIDE2IDE0NS42MSAxNS43MFExNDUuMDkgMTUuNDAgMTQ0LjQ1IDE1LjQwTDE0NC40NSAxNS40MFExNDMuODIgMTUuNDAgMTQzLjMwIDE1LjcwUTE0Mi43OSAxNiAxNDIuNDkgMTYuNTRRMTQyLjE5IDE3LjA5IDE0Mi4xOSAxNy44MFoiIGZpbGw9IiNGRkZGRkYiIHg9IjEyNi4zMSIvPjwvc3ZnPg==)](https://forthebadge.com)

Сервис для работы с заметками пользователей.

Используемые технологии:   
* PostgreSQL (В качестве хранилища данных)
* Docker (Для запуска сервиса)
* Swagger (Для документации API)
* Chi (маршрутизатор для создания HTTP-сервисов)
* cleanenv (Для конфигурации)
* pgx (драйвер для работы с PostgreSQL)

---

# Getting Started   

Для запуска нужно указать путь к конфигу в переменную окружения: CONFIG_PATH

---

# Usage

---

## Examples