import os

# Obtiene la ruta absoluta del directorio donde se encuentra el script
script_directory = os.path.dirname(os.path.abspath(__file__))

# Itera sobre todos los archivos en el directorio
for filename in os.listdir(script_directory):
    # Verifica si el archivo es un archivo .txt
    if filename.endswith(".TXT") or filename.endswith(".csv") or filename.endswith(".txt"):
        # Construye la ruta completa del archivo
    
        file_path = os.path.join(script_directory, filename)
        # Intenta eliminar el archivo
        try:
            os.remove(file_path)
            print(f"Archivo eliminado: {file_path}")
        except OSError as e:
            print(f"No se pudo eliminar {file_path}: {e}")