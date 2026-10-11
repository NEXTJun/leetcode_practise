# LeetCode 專案文件

本專案使用 Go 練習 LeetCode，重點是獨立推導、實作與面試口述。共用文件集中在 `docs/`；根目錄 `AGENTS.md` 是 Codex 的專案規則入口。

## 文件導覽

| 文件 | 用途 |
|---|---|
| [AGENTS.md](../AGENTS.md) | 固定教學原則與操作邊界 |
| [SESSION_HANDOFF.md](SESSION_HANDOFF.md) | 新對話接續用的當前狀態與下一步 |
| [PRACTICE_LOG.md](PRACTICE_LOG.md) | 按日期記練習摘要、久別診斷與是否需重練基礎 |
| [RECORD.md](RECORD.md) | 題目歷史與進度唯一總表 |
| [WORKFLOW.md](WORKFLOW.md) | 指令、保存、切題與批改流程 |
| [TEMPLATE.md](TEMPLATE.md) | 每題歸檔格式 |
| [RECENT_10_REVIEW.md](RECENT_10_REVIEW.md) | 最近十題複習快照（截至 2026-06-07 的紀錄） |

## 目錄

```text
.
├── AGENTS.md
├── main.go
├── docs/
│   ├── README.md
│   ├── WORKFLOW.md
│   ├── RECORD.md
│   ├── SESSION_HANDOFF.md
│   ├── PRACTICE_LOG.md
│   ├── TEMPLATE.md
│   └── RECENT_10_REVIEW.md
└── solved/
    └── <四位題號>_<英文_slug>/
        ├── my_attempt.md
        ├── summary.md
        └── final_solution.go
```

`solved/` 保留每題原始作答與筆記，方便與程式碼一起複習。若日後需要正式測試，再建立 `problems/` 與適用的 Go module。

## 執行與驗證

從專案根目錄執行當前作答：

```sh
go run main.go
```

目前沒有 `go.mod`；`solved/` 是複習用歸檔，可能依賴 LeetCode 提供的型別或 API。因此 `go test ./...` 目前不適合作為全專案通過標準。建立正式 module 與測試後，在該 module 根目錄執行：

```sh
go test ./...
```

若 Go cache 權限受限，可改用專案內快取：

```powershell
$env:GOCACHE = Join-Path (Get-Location) '.gocache'
go run main.go
```

```sh
GOCACHE="$PWD/.gocache" go run main.go
```

只有文件變更時，檢查路徑、連結、進度一致性與原始作答是否保留即可。

## 本次指令整理的依據

依 OpenAI 官方 [Rethinking skills and prompts for GPT-6 Astra](https://developers.openai.com/blog/rethinking-skills-and-prompts-for-gpt-6-astra)（查閱：2026-09-28），精簡重複指令、按任務讀取文件，並明確定義完成條件與操作界線。

本專案據此將固定規則、操作流程、當前狀態與歷史分開維護；保留逐步提示與原始作答保護。規則以教學需求為核心，不綁定模型名稱或更改 Codex 模型設定。
