@echo OFF
cd "runs"
del *.*

cd ".."



cd "./nongeneric"
FOR /L %%y IN (1, 1, 100) DO (
     go run main.go > NUL
    echo "non run number %%y"
)

cd "../generic"
FOR /L %%y IN (1, 1, 100) DO (
    go run main.go > NUL
    echo "generic run number %%y"
)

cd ".."
python results.py

PAUSE