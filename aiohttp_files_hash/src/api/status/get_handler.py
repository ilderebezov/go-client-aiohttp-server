from aiohttp import web
from pathlib import Path
from src.method.get_dir_list_files import get_dir_files_info
from src.method.get_dir_list_files import get_file_out


async def get_list(request: web.Request) -> web.Response:
    """
    curl -X GET http://localhost:8080/list
    :param request:
    :return:
    """

    await request.text()
    files = get_dir_files_info()
    return web.json_response(files, status=228)


async def get_file(request: web.Request) -> web.Response:
    """
    curl -X GET http://localhost:8080/file?name=xyz
    :param request:
    :return:
    """

    file_name_get = request.rel_url.query['name']
    if get_file_out(file_name_get):
        dir_path = Path.cwd()
        path = Path(dir_path, 'src', 'tmp', file_name_get)
        #path = Path(dir_path, 'tmp', file_name_get)
        return web.FileResponse(path, status=228)
    else:
        return web.Response(text='Error, the file is absent', status=227)
