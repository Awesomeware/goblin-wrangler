defmodule Kingpin.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  def start(_type, _args) do
    children = [
      # Start the Ecto repository
      Kingpin.Repo,
      # Start the Telemetry supervisor
      KingpinWeb.Telemetry,
      # Start the PubSub system
      {Phoenix.PubSub, name: Kingpin.PubSub},
      # Start the Endpoint (http/https)
      KingpinWeb.Endpoint
      # Start a worker by calling: Kingpin.Worker.start_link(arg)
      # {Kingpin.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Kingpin.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  def config_change(changed, _new, removed) do
    KingpinWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
