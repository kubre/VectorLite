# This is python file to generate some test embeddings to use

from sentence_transformers import SentenceTransformer

from numpy import ndarray
import numpy as np
from typing import Type

sentences = [
    "Chandrayaan-3 is the third Indian lunar exploration mission under the Indian Space Research Organisation's (ISRO) Chandrayaan programme.",
    "Chandrayaan-3 was launched on 14 July 2023.",
    "On 22 July 2019, ISRO launched Chandrayaan-2",
    "The Vikram lander is responsible for the soft landing on the Moon.",
    "The Pragyan rover is a six-wheeled vehicle with a mass of 26 kilograms (57 pounds).",
    "Confirming the existence of the project, ISRO's former chairman K. Sivan stated that the estimated cost would be around ₹615 crore (equivalent to ₹721 crore or US$90 million in 2023).",
]

model = SentenceTransformer("sentence-transformers/all-MiniLM-L6-v2")
embeddings: Type[ndarray] = model.encode(sentences, convert_to_numpy=True)

print("Generated Embeddings from sentence \n", embeddings)

np.asarray(embeddings)
np.savetxt("test_data.csv", embeddings, delimiter=",")
