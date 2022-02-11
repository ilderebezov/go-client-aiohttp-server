import aiohttp_cors
from aiohttp.web_app import Application

from src.api.status.get_handler import get_list
from src.api.status.get_handler import get_file
from src.api.status.put_handler import put_file
from src.api.status.post_file import post_file
from src.api.status.del_handler import delete_file


def setup_routes(app: Application):
    """Настраивает эндпоинты сервиса с поддержкой CORS."""
    cors = aiohttp_cors.setup(app, defaults={
        '*': aiohttp_cors.ResourceOptions(
            allow_credentials=True,
            expose_headers='*',
            allow_headers='*',
        ),
    })

    cors.add(app.router.add_get('/list', get_list))
    cors.add(app.router.add_get('/file', get_file))
    cors.add(app.router.add_put('/upload/file', put_file))
    cors.add(app.router.add_post('/update/file', post_file))
    cors.add(app.router.add_delete('/delete/file', delete_file))
