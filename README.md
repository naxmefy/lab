# lab

a simple to manage labs (simple monorepos)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/naxmefy/lab)
![Test](https://github.com/naxmefy/lab/workflows/Test/badge.svg)

## base structure of a lab

```base
lab_root
├───00001_hello_world
├───00002_hello_cargo
│   └───src
├───00003_unit_tests
│   └───src
└───...
```

## commands

| command | description                                                                                     |  state   |
| :------ | :---------------------------------------------------------------------------------------------- | :------: |
| create  | creates a new folder for the lab with an example for the given language                         | progress |
| add     | add the next lab subfolder (incremeting - e.g. if `00015_...` exists it will create `00016...`) | progress |
| remove  | removes the given folder index (will ask fo decrementing all following and confirmation)        | progress |
| exec    | executes the given command in all lab subfolder                                                 | progress |
