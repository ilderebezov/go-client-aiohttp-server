import os
from pathlib import Path
import hashlib


def get_file_out(file_name):
    """

    :param file_name:
    :return:
    """
    files_list = get_dir_files()
    if file_name in files_list:
        return True
    return False


def get_dir_files():
    """

    :return:
    """
    dir_path = Path.cwd()
    path = Path(dir_path, 'src', 'tmp')
    #path = Path(dir_path, 'tmp')
    files_list = os.listdir(path)
    files_list.remove('tmp_files')
    return files_list


def md5(file_name):
    """

    :param file_name:
    :return:
    """
    dir_path = Path.cwd()
    path = Path(dir_path, 'src', 'tmp', file_name)
    #path = Path(dir_path, 'tmp', file_name)
    hash_md5 = hashlib.md5()
    with open(path, "rb") as file_open:
        for chunk in iter(lambda: file_open.read(4096), b""):
            hash_md5.update(chunk)
    file_open.close()
    return hash_md5.hexdigest()


def get_dir_files_info():
    """

    :return:
    """
    files_list = get_dir_files()
    out_list = [{file_name: md5(file_name)} for file_name in files_list]
    return out_list
