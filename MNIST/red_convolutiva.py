import os
import logging
import time
from tensorflow.keras import models
from tensorflow.keras import layers
from tensorflow.keras.utils import to_categorical
import mnist_db as mnist
import numpy as np

logging.basicConfig(level=logging.INFO)

# Descargar datos si no existen
if not os.path.isdir("./data/"):
    logging.info("El directorio de datos no existe, creando...")
    os.mkdir("./data/")
    mnist.descargar_mnist('./data/')
elif not os.listdir('./data/'):
    logging.info("El directorio de datos está vacío")
    mnist.descargar_mnist('./data/')
else:
    logging.info("Los datos ya están descargados")

(training_images, training_labels, test_images, test_labels) = mnist.leer_mnist("./data/")

# Codificar las etiquetas
training_labels = to_categorical(training_labels)
test_labels = to_categorical(test_labels)

# Definir la topología de la red neuronal
network = models.Sequential()
network.add(layers.Conv2D(32, kernel_size=(3, 3), activation='relu', input_shape=(28, 28, 1)))
network.add(layers.MaxPooling2D((2, 2)))
network.add(layers.Conv2D(64, kernel_size=(3, 3), activation='relu'))
network.add(layers.MaxPooling2D((2, 2)))
network.add(layers.Flatten())
network.add(layers.Dense(512, activation='relu', kernel_initializer='random_uniform'))
network.add(layers.Dense(10, activation='softmax'))

# Preparar red especificando: función de error, optimizador y las métricas para evaluar su funcionamiento
network.compile(optimizer='rmsprop', loss='categorical_crossentropy', metrics=['accuracy'])

# Entrenar la red midiendo el tiempo que se tarda
train_start = time.time()
network.fit(training_images, training_labels, epochs=10, batch_size=128)
train_end = time.time()
print("Tiempo empleado en el entrenamiento: ", train_end - train_start)

# Evaluar la red sobre el conjunto de prueba y entrenamiento
print("Evaluando red neuronal...")
_, train_acc = network.evaluate(training_images, training_labels)
_, test_acc = network.evaluate(test_images, test_labels)
labels_prueba = network.predict(test_images)
print("Error sobre el conjunto de entrenamiento: ", 100.0 - (train_acc * 100.0), "%")
print("Error sobre el conjunto de prueba: ", 100.0 - (test_acc * 100.0), "%")

# Guardar etiquetas asignadas a los casos del conjunto de prueba en un fichero txt
f = open('labels-convol.txt', 'w')
for label in labels_prueba:
    print(np.argmax(label), end='', file=f)

f.close()