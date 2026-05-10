# kite-mcp-usecases

[![Go Reference](https://pkg.go.dev/badge/github.com/algo2go/kite-mcp-usecases.svg)](https://pkg.go.dev/github.com/algo2go/kite-mcp-usecases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Application use case layer for the algo2go ecosystem. Implements
write-side CQRS commands + read-side queries across all major
trading domains: orders (place/modify/cancel/close), portfolio
(holdings, P&L, dividends, sectors, returns), alerts (price,
composite, native), oauth bridge, sessions, tickers, options
strategy, paper trading, family accounts, data export, telegram
events, watchlists, consent, observability, plus the Saga + Ports
contracts.

Used by [`Sundeepg98/kite-mcp-server`](https://github.com/Sundeepg98/kite-mcp-server)
as the contract layer between the kc/manager state machine and the
MCP tool layer; every user-visible MCP tool dispatches through one
or more use cases here.

## Why a separate module?

The use case layer is the most user-visible surface of the
trading-platform domain — externalizing it completes the "ports &
adapters" architecture migration. Hosting as its own module:

- Centralizes write-side commands + read-side queries
- Lets command/query signatures version independently from the
  monolith
- Decouples saga orchestration from any one runtime
- Provides a stable contract for any consumer wiring algo2go
  trading domain primitives (alerts/orders/portfolio/etc.) into a
  custom MCP server, REST API, or direct embedding

## Stability promise

**v0.x — unstable.** Pin `v0.1.0` deliberately.

## Install

```bash
go get github.com/algo2go/kite-mcp-usecases@v0.1.0
```

## Public API (high-level)

- **Order use cases** — PlaceOrder, ModifyOrder, CancelOrder,
  ClosePosition, CloseAllPositions, ConvertPosition, GetOrders,
  pretrade check
- **Portfolio use cases** — GetPortfolio, P&L, options strategy
- **Alert use cases** — CreateAlert, CreateCompositeAlert,
  TrailingStop, NativeAlert
- **Account use cases** — Account, Admin, Family, OAuthBridge,
  Session, Setup, Consent, DataExport, Margin
- **Domain use cases** — Watchlist, GTT, MF (mutual funds), Ticker,
  Telegram events, Widget, Native alerts
- **Paper trading** — virtual portfolio mode use cases
- **Cross-cutting** — Observability, Pretrade, Saga, Ports
  contracts

## Dependencies (11 algo2go modules)

- `github.com/algo2go/kite-mcp-alerts` v0.1.0
- `github.com/algo2go/kite-mcp-broker` v0.1.0
- `github.com/algo2go/kite-mcp-cqrs` v0.1.0
- `github.com/algo2go/kite-mcp-domain` v0.1.0
- `github.com/algo2go/kite-mcp-eventsourcing` v0.1.0
- `github.com/algo2go/kite-mcp-logger` v0.1.0
- `github.com/algo2go/kite-mcp-money` v0.1.0
- `github.com/algo2go/kite-mcp-riskguard` v0.1.0
- `github.com/algo2go/kite-mcp-ticker` v0.1.0
- `github.com/algo2go/kite-mcp-users` v0.1.0
- `github.com/algo2go/kite-mcp-watchlist` v0.1.0
- `github.com/stretchr/testify` v1.10.0

All algo2go deps published; no upstream `replace` directives needed.

## Reference consumer

[`Sundeepg98/kite-mcp-server`](https://github.com/Sundeepg98/kite-mcp-server)
— consumed across 38 .go files: kc/manager_*, kc/ops/payoff.go,
app/*, mcp/* (admin, analytics, common, helpers, paper, plugin
widgets, portfolio, trade, watchlist, tax tools, tools_ext_apps).

## License

MIT — see [LICENSE](LICENSE).

## Authors

Original design: [Sundeepg98](https://github.com/Sundeepg98) (Zerodha
Tech). Multi-module promotion (2026-05-10): algo2go contributors.
