import os
import sys
import logging
from tensorflow.keras import models
from tensorflow.keras import layers
from tensorflow.keras.utils import to_categorical
import mnist_db as mnist
import numpy as np

logging.basicConfig(level=logging.INFO)

if len(sys.argv) == 1:
    print("Error, se debe indicar como parámetro el tipo de perceptron: single/multi")
    exit(-1)

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
# Una capa oculta de unidades lineales rectificadas y una capa de salida softmax
network = models.Sequential()

if sys.argv[1] == "single":
    network.add(layers.Input(shape=(28 * 28, )))
    network.add(layers.Dense(10, activation='softmax'))
elif sys.argv[1] == "multi":
    network.add(layers.Input(shape=(28 * 28,)))
    network.add(layers.Dense(512, activation='relu', input_shape=(28 * 28,), kernel_initializer='random_uniform'))
    network.add(layers.Dense(10, activation='softmax'))
else:
    print("Error, parámetro incorrecto")
    exit(-1)

# Preparar red especificando: función de error, optimizador y las métricas para evaluar su funcionamiento
network.compile(optimizer='rmsprop', loss='categorical_crossentropy', metrics=['accuracy'])

# Entrenar la red
network.fit(training_images, training_labels, epochs=10, batch_size=128)

# Evaluar la red sobre el conjunto de prueba
print("Evaluando red neuronal...")
_, test_acc = network.evaluate(test_images, test_labels)
labels_prueba = network.predict(test_images)
print("Precisión sobre el conjunto de prueba: ", test_acc * 100.0, "%")
print("Error sobre el conjunto de prueba: ", 100.0 - (test_acc * 100.0), "%")

# Guardar etiquetas asignadas a los casos del conjunto de prueba en un fichero txt
f = open('labels-' + sys.argv[1] + '.txt', 'w')
for label in labels_prueba:
    print(np.argmax(label), end='', file=f)

f.close()
