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
      KingpinWeb.Presence,
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

  def migrate do
    load_app()

    for repo <- repos() do
      :ok = ensure_repo_created(repo)
      {:ok, _, _} = Ecto.Migrator.with_repo(repo, &Ecto.Migrator.run(&1, :up, all: true))
    end
  end

  def rollback(repo, version) do
    load_app()
    {:ok, _, _} = Ecto.Migrator.with_repo(repo, &Ecto.Migrator.run(&1, :down, to: version))
  end

  defp ensure_repo_created(repo) do
    IO.puts "create #{inspect repo} database if it doesn't exist"
    case repo.__adapter__.storage_up(repo.config) do
      :ok -> :ok
      {:error, :already_up} -> :ok
      {:error, term} -> {:error, term}
    end
  end

  defp repos do
    Application.fetch_env!(:kingpin, :ecto_repos)
  end

  defp load_app do
    Application.load(:kingpin)
  end
end
