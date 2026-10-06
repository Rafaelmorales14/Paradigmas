### 1 - Python
a)
```
def adicionar(item, lista=[]):
    lista.append(item)
    return lista

print(adicionar(1))
print(adicionar(2))

Saída

[1]
[1, 2]
```
b) Passagem por atribuição.

### 2 - Java
a)
```
import java.util.*;

public class Main {
  static void zera(int[] v, int n) {
        v[0] = 0;
        n = 0;
        }

    public static void main(String[] args) {
      int[] v = {5, 5}; int n = 5;
      zera(v, n);
      System.out.println(v[0] + " " + n);
      }
}

Saída

0 5
```
b) Passagem por valor

### 3 - Python
a)
```
fs = [lambda: i for i in range(3)]
print([f() for f in fs])

Saída

[2, 2, 2]
```
b) Chamada indireta: lambda

 ### 4 - C

 a)

```
#include <stdio.h>

int contador(void) {
    static int n = 0;
    return ++n;
}

int main(){
    contador(); contador();
    printf("%d\n", contador());
}

Saída

3
```

 b) Variável local estática.

 ### 5 - Rust

 a)

```
fn dobra(v: Vec<i32>) -> Vec<i32> {
    v.iter().map(|x| x * 2).collect()
}

fn main() {
    let v = vec![1, 2, 3];
    let d = dobra(v);
    println!("{:?} {:?}", v, d);
}

Saída

error[E0382]: borrow of moved value: `v`
```

 b) Semântica de movimentação.

 ### 6 - Python

 a)

```
total = 0

def adiciona(x):
    total = total + x
    return total

print(adiciona(5))

Saída

UnboundLocalError: cannot access local variable 'total' where it is not associated with a value
```

 b) Escopo de variáveis.
