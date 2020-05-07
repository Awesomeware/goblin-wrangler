FROM alpine:3.11 as build

RUN apk add --no-cache elixir npm

RUN mix local.hex --force && mix local.rebar --force

COPY . .

RUN MIX_ENV=prod mix ci

FROM alpine:3.11

RUN apk add --no-cache ncurses-libs

COPY --from=build _build/prod/rel/kingpin /app

ENTRYPOINT ["/app/bin/kingpin"]
CMD ["start"]