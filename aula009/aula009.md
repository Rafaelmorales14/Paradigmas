### 1 - Java

a)

```
class Animal {
  String som() { return "..."; }
}
class Cachorro extends Animal {
  String som() { return "au"; }
}

public class Main {
    public static void main(String[] args) {
      Animal x = new Cachorro();
      System.out.println(x.som());
    }
}


Saída

au
```

b) Execução.

### 2 - C++

a)

```
#include <iostream>
using namespace std;

struct A {
  void f() { cout << "A"; }
};

struct B : A {
  void f() { cout << "B"; }
};

int main()
{
    B b; A *p = &b;
    p->f();
}

Saída

A
```

b) Compilação.

### 3 - Java

a)

```
import java.util.*;

class A { String nome = "A";
  String getNome() { return nome; } }
class B extends A { String nome = "B";
  String getNome() { return nome; } }

public class Main {
    public static void main(String[] args) {
      A x = new B();
      System.out.println(x.nome + " " + x.getNome());
    }
}

Saída:

A B
```

b) Compilação e Execução.


### 4 - Python

a)

```
class Contador:
    total = 0

    def __init__(self):
        Contador.total += 1
        self.id = Contador.total

a = Contador()
b = Contador()

print(a.id, b.id, a.total)

Saída

1 2 2
```

b) Atributo de classe e resolução de atributos em tempo de execução.

### 5 - Java

a)

```
class A {
    static String quem() {
        return "A";
    }
}

class B extends A {
    static String quem() {
        return "B";
    }
}

A x = new B();

System.out.println(x.quem());

Saída

A
```

b) Métodos `static`: decisão em tempo de compilação.

### 6 - Go

a)

```
type Animal struct{}

func (Animal) Som() string {
    return "..."
}

func (a Animal) Falar() string {
    return "faz " + a.Som()
}

type Cao struct{ Animal }

func (Cao) Som() string {
    return "au"
}

fmt.Println(Cao{}.Falar(), Cao{}.Som())

Saída

faz ... au
```

b) Composição por embedding e resolução estática de métodos.