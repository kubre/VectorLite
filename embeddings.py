# This is python file to generate some test embeddings to use

from sentence_transformers import SentenceTransformer

from numpy import ndarray
import numpy as np
from typing import Type


def generate_test_data(sentences: list[str], filename: str):
    model = SentenceTransformer("sentence-transformers/all-MiniLM-L6-v2")
    embeddings: Type[ndarray] = model.encode(sentences, convert_to_numpy=True)

    np.asarray(embeddings, np.float64)
    np.savetxt(f"{filename}.csv", embeddings, delimiter=",")


def cosine_similarity(a: np.array, b: np.array):
    return np.dot(a, b) / (np.linalg.norm(a) * np.linalg.norm(b))


if __name__ == "__main__":
    dataset = [
        "Chandrayaan-3 is the third Indian lunar exploration mission under the Indian Space Research Organisation's (ISRO) Chandrayaan programme.",
        "Chandrayaan-3 was launched on 14 July 2023.",
        "On 22 July 2019, ISRO launched Chandrayaan-2",
        "The Vikram lander is responsible for the soft landing on the Moon.",
        "The Pragyan rover is a six-wheeled vehicle with a mass of 26 kilograms (57 pounds).",
        "Confirming the existence of the project, ISRO's former chairman K. Sivan stated that the estimated cost would be around ₹615 crore (equivalent to ₹721 crore or US$90 million in 2023).",
    ]

    questions = ["Name of the space programme", "Number of wheels of vehicle"]

    # generate_test_data(dataset, "dataset")

    dataset = np.loadtxt("dataset.csv", dtype=np.float64, delimiter=",")
    questions = np.loadtxt("questions.csv", dtype=np.float64, delimiter=",")

    rankings = []
    for question in questions:
        temp = []
        for sentence in dataset:
            temp.append(cosine_similarity(question, sentence))
        rankings.append(temp)

    print(rankings)

""" Test Match using consine run
1.
[[0.33493956224613214, 0.2445450988872635, 0.2272637197659734, 0.13958924079453072, 0.07769994142277956, 0.12028321422009922],

[0.044401158939493944, 0.11715172866550795, 0.04526464701947782, 0.07302163296217688, 0.49793287036259076, 0.025503200589289506]]
"""
