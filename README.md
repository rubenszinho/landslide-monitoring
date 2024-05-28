# Sistema de Comunicação para Monitoramento de Desastres

## Objetivo do Projeto

Este projeto visa implementar um sistema de comunicação multiprotocolo para monitoramento de desastres, utilizando dispositivos ESP32 como unidades de comunicação e um coordenador central para gerenciar as comunicações. O sistema utiliza o protocolo MQTT para garantir uma comunicação confiável e eficiente entre as unidades e o coordenador.

---

## Coordenador

### Visão Geral

O coordenador é responsável por gerenciar a comunicação entre as diferentes unidades de campo (ESP32) e processar os dados recebidos. Ele recebe instruções de controle, envia sinais de "wake-up" para as unidades conforme necessário e encaminha os dados processados para o sistema externo.

### Tópicos MQTT

- **`/Control/Event/#`**: 
  - **Descrição**: O coordenador se inscreve nesse tópico para monitorar eventos de controle que possam exigir ações específicas, como instruções para acordar unidades ou mudar o estado do sistema.
  - **Racionalidade**: Este tópico permite que o coordenador responda a eventos globais que podem afetar o funcionamento do sistema como um todo.

- **`/Data/Coordinator/#`**:
  - **Descrição**: Tópico usado pelo coordenador para receber dados ou instruções específicas das unidades ou de outros sistemas de controle.
  - **Racionalidade**: Centraliza o recebimento de dados e instruções para o coordenador, facilitando o processamento e a disseminação de informações para as unidades.

- **`/Control/WakeUp/#`**:
  - **Descrição**: Tópico usado para enviar instruções de "wake-up" para as unidades que estão em estado de baixo consumo de energia.
  - **Racionalidade**: Garante que as unidades possam ser ativadas conforme necessário para processar eventos ou coletar dados, mantendo a eficiência energética.

- **`/Data/From/#`**:
  - **Descrição**: O coordenador se inscreve nesse tópico para receber dados enviados pelas unidades de campo.
  - **Racionalidade**: Permite que o coordenador agregue e processe os dados recebidos das unidades, encaminhando-os conforme necessário.

- **`/Data/To/Unit/{ID}`**:
  - **Descrição**: Tópico usado para enviar dados específicos para uma unidade identificada por seu ID.
  - **Racionalidade**: Facilita o envio de dados ou instruções direcionadas para unidades específicas, permitindo um controle granular das operações.

### Funcionamento do Coordenador

1. **Monitoramento de Eventos**:
   - O coordenador se inscreve em `/Control/Event/#` para monitorar eventos de controle.
   - Ao receber um evento, ele publica instruções correspondentes em `/Data/Coordinator/`.

2. **Recebimento e Processamento de Dados**:
   - O coordenador recebe dados em `/Data/Coordinator/#`.
   - Verifica se as unidades estão acordadas; se necessário, envia instruções de "wake-up" em `/Control/WakeUp/#`.
   - Encaminha os dados para as unidades apropriadas em `/Data/To/Unit/{ID}`.

3. **Wake-Up de Unidades**:
   - O coordenador monitora `/Control/WakeUp/#` para enviar sinais de ativação para as unidades.
   - As unidades, ao acordarem, podem enviar dados de volta através de `/Data/From/#`.

---

## Unidade de Comunicação (ESP32)

### Visão Geral

As unidades de comunicação são dispositivos ESP32 que se conectam ao coordenador via MQTT para receber instruções, coletar dados e transmiti-los ao coordenador ou a sistemas externos. Elas são capazes de entrar em estado de baixo consumo de energia e acordar sob demanda, conforme instruções do coordenador.

### Tópicos MQTT

- **`/Data/To/Unit/{ID}`**:
  - **Descrição**: Tópico onde a unidade se inscreve para receber dados ou instruções específicas do coordenador.
  - **Racionalidade**: Permite que a unidade receba dados ou comandos diretamente do coordenador, facilitando a execução de tarefas específicas.

- **`/Data/From/Unit/{ID}`**:
  - **Descrição**: Tópico onde a unidade publica dados processados para serem enviados de volta ao coordenador.
  - **Racionalidade**: Facilita o envio de dados coletados ou processados pela unidade, garantindo que o coordenador possa receber e processar essas informações.

### Funcionamento da Unidade de Comunicação

1. **Conexão e Subscrição**:
   - A unidade se conecta ao broker MQTT e se inscreve no tópico `/Data/To/Unit/{ID}` para receber dados e instruções.
   - Mantém uma sessão persistente e utiliza QoS 1 para garantir a entrega dos dados.

2. **Recepção de Dados**:
   - Ao receber uma mensagem no tópico `/Data/To/Unit/{ID}`, a unidade processa a mensagem e retransmite os dados para o sistema externo via `RelayData()`.

3. **Envio de Dados ao Coordenador**:
   - Após processar os dados, a unidade envia os resultados de volta ao coordenador publicando no tópico `/Data/From/Unit/{ID}`.

4. **Manutenção da Conexão**:
   - A unidade verifica continuamente a conexão com o broker MQTT e se reconecta automaticamente em caso de desconexão.
   - Utiliza QoS 1 para garantir que as mensagens sejam entregues pelo menos uma vez, garantindo a confiabilidade do sistema.
