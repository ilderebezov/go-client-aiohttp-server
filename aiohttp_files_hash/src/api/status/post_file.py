import os
from pathlib import Path

from aiohttp import web

from src.method.get_dir_list_files import get_dir_files
from src.method.get_dir_list_files import md5


async def post_file(request: web.Request) -> web.Response:
    """
    curl -F 'data=@test2' http://localhost:8080/update/file
    :param request:
    :return:
    """

    upload = await request.post()
    upload_data = upload.get("data")
    upload_data_content = upload_data.file.read().decode('utf-8')
    file_name = upload_data.filename

    files_list = get_dir_files()
    dir_path = Path.cwd()
    tmp_path = Path(dir_path, 'src', 'tmp', 'tmp_files', file_name)
    #tmp_path = Path(dir_path, 'tmp', 'tmp_files', file_name)
    with open(tmp_path, 'w', encoding='utf-8') as file_open:
        file_open.write(upload_data_content)
    file_open.close()
    upload_file_hash = md5(tmp_path)

    if file_name in files_list:
        path_file_server = Path(dir_path, 'src', 'tmp', file_name)
        #path_file_server = Path(dir_path, 'tmp', file_name)
        file_server_hash = md5(path_file_server)
        if upload_file_hash == file_server_hash:
            os.remove(tmp_path)
            return web.Response(text=f'The upload file with filename: '
                                     f'{file_name} has similar hash as '
                                     f'server file, nothing to update.', status=227)
        else:
            os.remove(tmp_path)
            with open(path_file_server, 'w', encoding='utf-8') as file_open:
                file_open.write(upload_data_content)
            file_open.close()
            return web.Response(text=f'The file with filename: '
                                     f'{file_name} is update on server.', status=229)
    else:
        return web.Response(text=f'Error, file with filename: '
                                 f'{file_name} is absent on server', status=228)
