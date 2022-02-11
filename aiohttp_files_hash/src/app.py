from aiohttp import web
from src.service import routers


def init_app() -> web.Application:
    """Инициализация приложения со всеми настройками."""
    app = web.Application()
    routers.setup_routes(app)
    return app


def start():
    app = init_app()
    web.run_app(app)


if __name__ == '__main__':
    start()
