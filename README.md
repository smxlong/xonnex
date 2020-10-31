<h1>xonnex - a graph library for Go</h1>

*Yeah, there are quite a few of these already, but here's mine.*

---

- [Introduction](#introduction)
  - [Graph elements](#graph-elements)
    - [Nodes](#nodes)
      - [Node data](#node-data)
      - [Node metadata](#node-metadata)
    - [Edges](#edges)

---

# Introduction

Xonnex is a library for graph processing. It provides data structures to
represent the different elments of a graph.

## Graph elements

### Nodes

Nodes are the vertices in a graph.

A node has associated *data* as well as *metadata*.

#### Node data

Nodes can hold data, in the form of `interface{}`.

#### Node metadata

Nodes have metadata. See

### Edges

Edges are the lines between vertices in a graph.

An edge is associated with two nodes called *from* and *to*. The edge connects
these two nodes. An edge can be *directional* or not. If an edge is directional,
there is a distinction between *from* and *to*. If the edge is not directional,
there is no such distinction.