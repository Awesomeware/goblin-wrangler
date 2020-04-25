defmodule KingpinWeb.PageController do
  use KingpinWeb, :controller

  def index(conn, _params) do
    render(conn, "index.html")
  end
end
