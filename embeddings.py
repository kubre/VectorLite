# This is python file to generate some test embeddings to use

from sentence_transformers import SentenceTransformer

import numpy as np
import csv


def generate_test_data(sentences: list[str], filename: str):
    model = SentenceTransformer("sentence-transformers/all-MiniLM-L6-v2")
    embeddings = model.encode(sentences, convert_to_numpy=True)
    new_embeddings = [
        [sentence, *embedding] for sentence, embedding in zip(sentences, embeddings)
    ]

    with open(filename, "w", encoding="utf-8", newline="") as f:
        writer = csv.writer(f)
        writer.writerows(new_embeddings)

    # np.savetxt(f"{filename}.csv", new_embeddings, delimiter=",")


def load_data_from_file(filename: str):
    temp = []
    with open(filename, "r", encoding="utf-8", newline="") as f:
        reader = csv.reader(f)
        for row in reader:
            temp2 = [row[0]]
            for value in row[1:]:
                temp2.append(float(value))
            temp.append(temp2)
    return temp


def cosine_similarity(a: np.array, b: np.array):
    return np.dot(a, b) / (np.linalg.norm(a) * np.linalg.norm(b))


if __name__ == "__main__":
    text_sentences = [
        "Chandrayaan-3 is the third Indian lunar exploration mission under the Indian Space Research Organisation's (ISRO) Chandrayaan programme.",
        "Chandrayaan-3 was launched on 14 July 2023.",
        "On 22 July 2019, ISRO launched Chandrayaan-2",
        "The Vikram lander is responsible for the soft landing on the Moon.",
        "The Pragyan rover is a six-wheeled vehicle with a mass of 26 kilograms (57 pounds).",
        "Confirming the existence of the project, ISRO's former chairman K. Sivan stated that the estimated cost would be around ₹615 crore (equivalent to ₹721 crore or US$90 million in 2023).",
    ]

    text_questions = ["Name of the space programme", "Number of wheels of vehicle"]

    # sentences = generate_test_data(text_sentences, "sentences.csv")
    # questions = generate_test_data(text_questions, "questions.csv")

    sentences = load_data_from_file("sentences.csv")
    questions = load_data_from_file("questions.csv")

    rankings = []
    for question in questions:
        temp = []
        for sentence in sentences:
            temp.append([sentence[0], cosine_similarity(question[1:], sentence[1:])])
        rankings.append(temp)

    for rank in rankings:
        data = sorted(rank, key=lambda x: x[1], reverse=True)
        print(data[0])


""" Test Match using consine run
1.
[[0.33493956224613214, 0.2445450988872635, 0.2272637197659734, 0.13958924079453072, 0.07769994142277956, 0.12028321422009922],

[0.044401158939493944, 0.11715172866550795, 0.04526464701947782, 0.07302163296217688, 0.49793287036259076, 0.025503200589289506]]
"""

""" Output from Go
[[0.33493956224613236 0.2445450988872636 0.22726371976597373 0.13958924079453083 0.07769994142277957 0.12028321422009919]

[0.04440115893949398 0.11715172866550799 0.04526464701947791 0.07302163296217701 0.49793287036259126 0.025503200589289534]]
"""
