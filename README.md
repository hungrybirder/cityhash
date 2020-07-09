# Go cityhash #

Go implementation of Google City Hash.

Clone from [zhenjl/cityhash](https://github.com/zhenjl/cityhash).


# Install Package #
```bash
go get github.com/hungrybirder/cityhash
```

# Install cityhash64 cli #
```bash
go get github.com/hungrybirder/cityhash/cmd/cityhash64@master

```

# Usage cli #
```bash
cityhash64 -seed=11 hello world "123 456"
==========Show CityHash64 Demos==========

hash                          	hash(hex)                     	key
3852287081031594632           	357612b09db6f288              	hello
3366074994012428358           	2eb6b35f716f9846              	world
239141937353951750            	3519a5332e18a06               	123 456

Seed=11
```

```bash
cityhash64 hello world "123 456"
==========Show CityHash64 Demos==========

hash                          	hash(hex)                     	key
13009744463427800296          	b48be5a931380ce8              	hello
16436542438370751598          	e41a54435eb8b46e              	world
8183505631039799034           	7191a7036b893efa              	123 456

```
