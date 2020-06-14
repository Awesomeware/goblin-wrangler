defmodule Kingpin.Groups.Group do
  @moduledoc "Groups are sets of users/players, who play in a game together"

  use Ecto.Schema
  import Ecto.Changeset

  schema "groups" do
    field :name, :string
    belongs_to :leader, Kingpin.Accounts.User
    many_to_many :users, Kingpin.Accounts.User, join_through: Kingpin.Groups.GroupUser, unique: true

    timestamps()
  end

  @doc false
  def changeset(user, attrs) do
    user
    |> cast(attrs, [:name, :leader_id])
    |> validate_required([:name, :leader_id])
  end
end

defmodule Kingpin.Groups.GroupUser do

  use Ecto.Schema
  import Ecto.Changeset

  @primary_key false
  schema "group_users" do
    belongs_to :user, Kingpin.Accounts.User
    belongs_to :group, Kingpin.Groups.Group

    timestamps()
  end

  def changeset(group_user, attrs \\ %{}) do
    group_user
    |> cast(attrs, [:user_id, :group_id])
    |> validate_required([:user_id, :group_id])
  end
end
