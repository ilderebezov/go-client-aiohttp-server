import os
from pathlib import Path

from aiohttp import web

from src.method.get_dir_list_files import get_dir_files


async def delete_file(request: web.Request) -> web.Response:
    """

    :param request:
    :return:
    """
    # curl -d 'test' -X DELETE http://localhost:8080/delete/file
    file_name_del = await request.text()
    files_list = get_dir_files()

    if file_name_del in files_list:
        dir_path = Path.cwd()
        path = Path(dir_path, 'src', 'tmp', file_name_del)
        #path = Path(dir_path, 'tmp', file_name_del)
        os.remove(path)
        return web.Response(text=f'The file with filename: '
                                 f'{file_name_del} is remove'
                                 f'from server successfully.', status=229)
    else:
        return web.Response(text=f'Error, the file with filename: '
                                 f'{file_name_del} is absent'
                                 f' on server.', status=228)
