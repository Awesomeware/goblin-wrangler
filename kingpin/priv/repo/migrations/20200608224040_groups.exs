defmodule Kingpin.Repo.Migrations.Groups do
  use Ecto.Migration

  def change do
    create table(:groups) do
      add :name, :string
      add :leader, references(:users)

      timestamps()
    end

    create table(:group_users) do
      add :user_id, references(:users)
      add :group_id, references(:groups)

      timestamps()
    end
  end
end
