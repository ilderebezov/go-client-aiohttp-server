from pathlib import Path

from aiohttp import web

from src.method.get_dir_list_files import get_dir_files


async def put_file(request: web.Request) -> web.Response:
    """
    curl -X PUT "localhost:8080/upload/file" -F "file=@test3" -b "JSESSIONID=cookievalue"
    :param request:
    :return:
    """

    upload = await request.post()
    upload_data = upload.get("file")
    upload_data_content = upload_data.file.read().decode('utf-8')
    file_name = upload_data.filename

    files_list = get_dir_files()
    if file_name in files_list:
        return web.Response(text=f'Error, the upload file with filename: '
                                 f'{file_name} is exist on server.', status=228)
    else:
        dir_path = Path.cwd()
        path_file_server = Path(dir_path, 'src', 'tmp', file_name)
        #path_file_server = Path(dir_path, 'tmp', file_name)
        with open(path_file_server, 'w', encoding='utf-8') as file_open:
            file_open.write(upload_data_content)
        file_open.close()
        return web.Response(text=f'Upload file with filename {file_name} on server is successful', status=229)

