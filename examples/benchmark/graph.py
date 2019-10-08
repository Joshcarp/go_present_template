import csv
import os
import pandas
import matplotlib.pyplot as plt
import numpy as np
from itertools import cycle
import re

def convertToCsv(filename):
    original = open(filename, "r")
    output = open("results.csv", "w")
    c = csv.writer(output)
    next(original)
    next(original)
    c.writerow(["TestName", "LibraryName", "iterations", "time"])
    line = original.read()
    regex = r"(?:BenchmarkDecimal\/dd)(?P<TestName>\w*)(?:.decTest_)(?P<LibraryName>\w*)(?:(#\d*)*-\d\s*)(?P<iterations>\d*)(?:\s*)(?P<time>\d*)"
    matches = re.finditer(regex, line)
    for mat in matches:
        c.writerow([mat.group("TestName"), mat.group("LibraryName"), mat.group("iterations"), int(mat.group("time"))/1000000])
    output.close()
    return





def generate_graph(responsetime_filename):
    data = pandas.read_csv(responsetime_filename)
    plt.rcParams.update({'font.size': 17})
    operations = data.TestName.unique()
    for operation in operations:
        element = data.loc[data["TestName"] == operation]
        duplicates = element.duplicated(subset = 'LibraryName', keep = 'first')
        dupes = element[duplicates]
        print(element)

        # [dupes.append(pandas.Series(), ignore_index=True) for i in range(len(element.index)-len(dupes.index))] 
        element = element[~duplicates]
        print(len(element[~duplicates].index))
        plt.figure(figsize=(10,6))
        x = np.arange(len(element.index))
        dupeInd = np.arange(len(dupes.index))
        bar_lst2 = []
        # bar_lst2 = plt.bar(dupeInd, dupes.time.tolist())
        bar_lst = plt.bar(x, element.time.tolist())
       
        colors = cycle([(232/255, 62/255, 93/255, 0.3),(65/255, 186/255, 99/255, 0.3),(87/255, 126/255, 199/255, 0.3),(232/255, 165/255, 0, 0.3)])
        for i in range(len(bar_lst)):
            bar_lst[i].set_color(next(colors))
        for i in range(len(bar_lst2)):
            next(colors)
            bar_lst2[i].set_color(next(colors))
        plt.xticks(x, element.LibraryName)
        plt.xticks(rotation=10)

        plt.ylabel('ns/operation')
        plt.title(operation + " Benchmark")
        plt.savefig(os.path.abspath('../../content/')+f"/img/{operation}_new.png", dpi=300, figsize=(10,25),  bbox_inches = 'tight')

if __name__ == "__main__":
    # os.system("go test -bench=. > results.txt")
    convertToCsv("responses.txt")
    generate_graph("results.csv")
