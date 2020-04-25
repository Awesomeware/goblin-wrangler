defmodule Kingpin.Repo do
  use Ecto.Repo,
    otp_app: :kingpin,
    adapter: Ecto.Adapters.Postgres
end
