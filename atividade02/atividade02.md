# Evolução das Principais Linguagens de Programação

## 1. A genealogia das linguagens não é uma escada de progresso. Explique essa afirmação e apresente dois fatores históricos que fazem uma linguagem influenciar outra sem necessariamente substituí-la.

<p>A história das linguagens não acontece de forma sequêncial, onde cada linguagem criada é melhor e deve substituir a anterior. Em vez disso, diferentes linguagens existem de forma simultânea e influenciando uma a outra, além de possuirem contextos, estruturas, paradigmas e objetivos diferentes.</p>

<p>Dois fatores históricos que explicam isso são:</p>

- Necessidades e contextos diferentes: uma linguagem pode ser criada para determinada finalidade, como sistemas embarcados, aplicações científicas ou desenvolvimento web, e continuar adequada mesmo depois do surgimento de linguagens mais novas. Exemplos: JavaScript (Navegadores) e Java (Multiplataforma).

- Legado e comunidade: uma linguagem já utilizada por muitas empresas e programadores acumula programas, bibliotecas, ferramentas e conhecimento. Esse investimento torna difícil abandoná-la completamente. Ademais, linguagens novas frequentemente utilizam conceitos e recursos das anteriores, criando influência sem substituição.

---

## 2. Por ser o primeiro projeto de uma linguagem de alto nível da história, provou que algoritmos complexos poderiam ser escritos antes de existirem computadores capazes de fazer a leitura deles. Exemplos de recursos disponíveis:

O **Plankalkül**, projeto de Konrad Zuse, antecipou diversos recursos que posteriormente se tornariam comuns nas linguagens de programação. Mesmo sem ter sido implementado em um computador na época, demonstrava que algoritmos complexos poderiam ser descritos de forma estruturada e independente da execução imediata.

* **Vetores e registros (estruturas de dados):** permitia representar dados complexos a partir de elementos menores, além de possibilitar o aninhamento de estruturas, como colocar um vetor dentro de um registro.

* **Sub-rotinas (reutilização de funções):** permitia organizar partes do algoritmo em blocos separados, que poderiam ser reutilizados, facilitando a estruturação de problemas grandes sem a necessidade de repetir código.

* **Laços de repetição:** apresentava mecanismos próprios para representar iterações, evitando depender exclusivamente do comando `GOTO`. Também utilizava uma instrução de desvio chamada **Fin**, contribuindo para uma forma mais estruturada de descrever a execução.

---

## 4. Explique por que o projeto Fortran precisou convencer programadores de que código traduzido podia competir com código de máquina escrito à mão. Relacione desempenho, custo de programação e adoção.

Até o surgimento do **Fortran**, era comum considerar que programar diretamente em código de máquina ou Assembly era necessário para obter o melhor desempenho. Existia, portanto, uma desconfiança de que uma linguagem de alto nível, traduzida por um compilador, produziria programas lentos.

O Fortran procurou solucionar esse problema oferecendo **alto desempenho sem exigir que o programador escrevesse todo o código manualmente**. Seus compiladores conseguiam gerar código eficiente para a época, especialmente em operações numéricas e científicas.

* **Desempenho:** o código gerado pelo compilador podia alcançar um desempenho competitivo com programas escritos manualmente em Assembly, reduzindo a principal desvantagem atribuída às linguagens de alto nível.

* **Custo de programação:** escrever programas em Fortran era muito mais simples e rápido do que trabalhar diretamente com código de máquina ou Assembly, principalmente em cálculos numéricos e operações com ponto flutuante.

* **Adoção:** essa combinação tornou o trade-off favorável: mesmo que algum desempenho pudesse ser perdido em determinados casos, a enorme redução no esforço de programação compensava essa diferença. Com isso, o Fortran ajudou a provar que linguagens de alto nível poderiam ser utilizadas em aplicações de alto desempenho.

---

## 6. ALGOL 60 teve grande importância histórica mesmo não sendo amplamente utilizada no ramo comercial. Explique por quê.

O **ALGOL 60** teve uma importância histórica muito maior do que sua adoção comercial. Apesar das limitações dos compiladores, dos problemas de compatibilidade com projetos existentes e da dificuldade de migração, a linguagem introduziu e consolidou conceitos que influenciaram muitas linguagens posteriores.

* **Recursividade:** permitiu que procedimentos chamassem a si mesmos, possibilitando soluções mais naturais para determinados problemas.

* **Blocos e escopo:** organizou o código em blocos e estabeleceu regras mais claras para o alcance das variáveis, contribuindo para a programação estruturada.

* **BNF (Backus-Naur Form):** utilizou uma notação formal para especificar a sintaxe da linguagem, facilitando uma descrição rigorosa da estrutura dos programas.

Esses conceitos foram posteriormente incorporados ou desenvolvidos em linguagens como **C, C++, Pascal e Java**. Portanto, o ALGOL 60 mostrou que uma linguagem pode ter grande influência histórica mesmo sem alcançar ampla adoção comercial.

---

## 7. COBOL foi desenhada para processamento comercial. Mostre como domínio e público influenciaram sua legibilidade, seus registros e sua relação com FLOW-MATIC.

<p>O domínio comercial e o público-alvo influenciaram diretamente o projeto do COBOL. Como a linguagem era destinada ao processamento de dados de empresas e a usuários próximos da área administrativa, buscou-se uma sintaxe legível e semelhante ao inglês, facilitando a compreensão por profissionais que não fossem especialistas em programação.</p>

- Legibilidade: COBOL utiliza palavras-chave e estruturas próximas da linguagem natural, como IF, MOVE, ADD e PERFORM, priorizando programas que pudessem ser lidos e mantidos por equipes comerciais.
- Registros: a linguagem foi projetada para trabalhar bem com grandes volumes de dados estruturados, especialmente registros de arquivos comerciais, permitindo definir campos e estruturas de dados de forma explícita.
- Relação com FLOW-MATIC: COBOL recebeu forte influência da FLOW-MATIC, desenvolvida por Grace Hopper. FLOW-MATIC já buscava uma programação mais próxima do inglês e voltada ao processamento de dados empresariais. COBOL aproveitou essas ideias e as ampliou, tornando-se uma linguagem mais padronizada e adequada às necessidades comerciais.

<table>
    <head>
        <th>Orientado a dados</th>
        <th>Orientado a objetos</th>
    <head>
    <body>
        <tr>
            <td>Foco nos dados e seu processamento.</td>
            <td>Foco em objetos que juntam dados + comportamentos.</td>
        </tr>
        <tr>
            <td>Os dados são geralmente manipulados por procedimentos.</td>
            <td>Os objetos possuem atributos (dados) e métodos (ações).</td>
        </tr>
    </body>
</table>

---

## 11. ALGOL 60 → Pascal → C

O **ALGOL 60** consolidou conceitos fundamentais da programação, como estruturas de controle, blocos, escopo de variáveis, tipos e procedimentos. Dessa forma, serviu como uma importante base conceitual para diversas linguagens posteriores.

A **Pascal** foi fortemente inspirada no ALGOL e manteve sua orientação para a programação estruturada. Além disso, trouxe uma preocupação maior com **tipagem, organização dos dados e estruturas de programação**, tornando-se especialmente importante para o ensino de conceitos fundamentais de programação.

A linguagem **C** também recebeu forte influência da tradição do ALGOL e da programação estruturada e imperativa. Entretanto, enfatizou ainda mais a **eficiência, manipulação de memória, tipos e controle da execução**, tornando-se uma importante base para várias linguagens posteriores.

O **Prolog** segue um caminho diferente. Enquanto ALGOL, Pascal e C são baseadas principalmente na ideia de descrever **como** o computador deve executar uma tarefa, Prolog utiliza a **programação lógica**. Em vez de indicar passo a passo como realizar uma operação, o programador declara **fatos, regras e consultas**, permitindo que o sistema realize inferências para encontrar uma resposta.

---

## 12. Modele em linguagem natural uma pequena base Prolog com dois fatos, uma regra e uma consulta. Explique por que isso representa programação lógica, não apenas armazenamento de dados.

<p>Porque o programa não apenas guarda valores lógicos, ele se utiliza de fatos, regras e consiltas para adiquirir uma nova informação, como no seguinte exemplo, onde a regra permite deduzir uma nova informação: que o gato é um bom caçador.</p>

### Dois fatos:

```
- O gato tem pelos.
- O gato caça ratos.
```

### Uma regra:

```
- Se um animal tem pelos e caça ratos, então ele é um bom caçador.
```

### Consulta:

```
- O gato é um bom caçador?
```

### Em Prolog:

```
tem_pelos(gato).
caca_ratos(gato).

bom_cacador(X) :- tem_pelos(X), caca_ratos(X).


?- bom_cacador(gato).
```
**Resultado: true**

---

## 13. Ada resultou de requisitos e projeto em grande escala. Analise como confiabilidade, tipos, pacotes e concorrência se relacionam ao domínio de sistemas críticos.

A linguagem Ada foi desenvolvida para atender às necessidades do Departamento de Defesa dos Estados Unidos, especialmente em sistemas embarcados e de grande escala. Esse domínio exigia uma linguagem que priorizasse confiabilidade, segurança e organização, já que falhas poderiam causar consequências graves.

- **Confiabilidade:** Ada possui mecanismos que ajudam a detectar erros ainda durante a compilação, aumentando a segurança e reduzindo falhas em tempo de execução.

- **Tipos:** sua tipagem forte permite definir dados de maneira precisa e restringir operações inadequadas, ajudando o compilador a identificar erros antes da execução.

- **Pacotes:** permitem organizar o código em módulos, separando interfaces e implementações. Isso facilita a manutenção e o desenvolvimento de sistemas grandes e complexos.

- **Concorrência:** Ada possui recursos próprios para concorrência, permitindo que diferentes tarefas sejam executadas e coordenadas simultaneamente, algo importante em sistemas embarcados que precisam responder a vários eventos ao mesmo tempo.

Assim, essas características estão diretamente relacionadas às exigências de sistemas críticos, nos quais segurança, previsibilidade e confiabilidade são fundamentais.

---

## 15. A primeira aplicação de Java não foi a Web, mas a Web impulsionou sua adoção. Explique como mudanças de contexto podem reposicionar uma linguagem.

O **Java** foi inicialmente desenvolvido com foco em dispositivos e sistemas embarcados, dentro do projeto que posteriormente deu origem à linguagem. Entretanto, o contexto original não alcançou o sucesso comercial esperado.

Com a expansão da **Web**, surgiu uma nova oportunidade para a linguagem. O Java apresentava características que se encaixavam bem nesse novo contexto, principalmente a possibilidade de executar o mesmo programa em diferentes plataformas por meio da **máquina virtual Java (JVM)**. Os applets também contribuíram inicialmente para sua popularização na Web.

Dessa forma, uma mudança no contexto tecnológico fez com que uma linguagem criada para uma finalidade pudesse encontrar uma aplicação muito mais relevante em outra. O Java foi reposicionado e passou a ter grande importância no desenvolvimento de aplicações Web e, posteriormente, em diversos outros tipos de sistemas.

---

## 17. C# foi apresentada como evolução no ambiente .NET. Compare duas decisões de C# com suas correspondentes em Java ou C++ e explique o problema que pretendem resolver.

Uma diferença entre **C# e Java** está no tratamento do comando `goto`. O Java não permite seu uso, enquanto o C# mantém essa possibilidade, assim como ocorre em C e C++. Embora o `goto` possa tornar o código mais difícil de compreender quando utilizado de forma excessiva, sua presença oferece maior flexibilidade para determinados casos de controle de fluxo.

Outra diferença aparece nos **enums**. O C# possui uma implementação de enum mais controlada e segura do que a encontrada tradicionalmente em C++, pois não permite simplesmente tratar qualquer valor de enum como um `int` sem conversão. Nesse aspecto, aproxima-se da preocupação do Java com maior segurança de tipos.

Essas decisões mostram como o C# buscou combinar **flexibilidade herdada de linguagens como C++** com **maior segurança e organização**, características associadas ao Java e ao ambiente .NET.
