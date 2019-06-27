import csv
import os

import pandas
import matplotlib.pyplot as plt
from matplotlib.ticker import FuncFormatter
import matplotlib.pyplot as plt
import numpy as np
from itertools import *
import re

def convertToCsv(filename):
    original = open(filename, "r")
    output = open("results.csv", "w")
    c = csv.writer(output)
    next(original)
    next(original)
    c.writerow(["TestName", "LibraryName", "iterations", "time"])
    # original = list(original)[2:-2]
    line = original.read()
    regex = r"(?:BenchmarkDecimal\/dd)(?P<TestName>\w*)(?:.decTest_)(?P<LibraryName>\w*)(?:-\d\s*)(?P<iterations>\d*)(?:\s*)(?P<time>\d*)"
    matches = re.finditer(regex, line)
    for mat in matches:
        c.writerow([mat.group("TestName"), mat.group("LibraryName"), mat.group("iterations"), mat.group("time")])
    output.close()
    return





def generate_graph(responsetime_filename):
    data = pandas.read_csv(responsetime_filename)
    for element in data:
        print(element)


    Add = data.loc[data["TestName"] == "Add"]
    Multiply = data.loc[data["TestName"] == "Multiply"]
    Abs = data.loc[data["TestName"] == "Abs"]
    Divide = data.loc[data["TestName"] == "Divide"]
    Graphs = [Add, Multiply, Abs, Divide]

    for element in Graphs:
        print(element)
        name = element.iloc[1]["TestName"]
        plt.figure(figsize=(10,6))
        x = np.arange(len(element.index))
        bar_lst = plt.bar(x, element.time.tolist())
        colors = cycle([(232/255, 62/255, 93/255, 0.3),(65/255, 186/255, 99/255, 0.3),(87/255, 126/255, 199/255, 0.3),(232/255, 165/255, 0, 0.3)])
        for i in range(len(bar_lst)):
            bar_lst[i].set_color(next(colors))
        plt.xticks(x, element.LibraryName)
        plt.xticks(rotation=10)
        plt.title(name + " Benchmark")
        # plt.ylabel('Time taken (ns)')
        plt.savefig(os.path.abspath('../../content/')+f"/img/{name}.png", dpi=300, figsize=(50,25))

# os.system("go test -bench=. > results.txt")
c = convertToCsv("results.txt")

generate_graph("results.csv")
