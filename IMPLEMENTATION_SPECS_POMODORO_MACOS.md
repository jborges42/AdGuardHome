# Especificação de Implementação — App Pomodoro + Bloqueio + Microlearning (macOS)

## 0) Objetivo deste documento
Traduzir o plano de produto em uma especificação executável de implementação para ambiente macOS, com foco em:
- comportamento de cada feature;
- requisitos funcionais e não funcionais;
- arquitetura alvo e fronteiras de responsabilidade;
- riscos técnicos e decisões de engenharia;
- critérios de prontidão por fase.

> Escopo: definição arquitetural e comportamental. Sem detalhamento de código.

---

## 1) Princípios arquiteturais (decisões de base)

## 1.1 Premissas
1. O cliente principal é macOS e deve parecer nativo.
2. Bloqueio de apps depende de permissões do sistema (Accessibility, Screen Recording, etc.).
3. Microlearning com IA exige pipeline com confiabilidade rastreável.
4. Produto deve funcionar offline de forma degradada (timer, bloqueio, histórico local).

## 1.2 Decisões estruturantes
- DA-01: **Cliente nativo macOS** (SwiftUI + integração AppKit quando necessário).
- DA-02: **Domínio desacoplado de UI** (camadas: Presentation, Application, Domain, Infrastructure).
- DA-03: **Serviço de conteúdo IA separado** (local service ou backend remoto), para permitir evolução sem acoplar ao app.
- DA-04: **Persistência local obrigatória** para estado de sessão, trilha e relatório.
- DA-05: **Telemetria mínima orientada a produto** para medir aderência, qualidade e risco operacional.

## 1.3 Critérios de decisão técnica (sempre validar antes de implementar)
- Impacta UX nativa do macOS?
- Exige permissão sensível? Qual fallback existe?
- A falha degrada para modo seguro?
- Métrica/KPI desta feature está instrumentada?
- Complexidade adicionada é proporcional ao valor entregue?

---

## 2) Arquitetura lógica (visão de alto nível)

## 2.1 Módulos principais
1. **Session Engine**
   - Orquestra ciclo Pomodoro (foco/pausa/pausa longa).
   - Controla estado, transições e timers.
2. **Block Policy Engine**
   - Aplica modo bloqueio por app ou bloqueio total (políticas de restrição).
   - Gerencia bypass e auditoria.
3. **Learning Engine**
   - Gerencia trilha (roadmap), progresso e seleção da próxima microlição.
4. **Content Orchestrator**
   - Busca/valida/sintetiza conteúdo com score de confiança.
5. **Report Engine**
   - Consolida tempo por assunto/categoria e histórico de sessão.
6. **Permissions Manager**
   - Avalia permissões de sistema e orienta usuário com UX clara.

## 2.2 Estados globais da aplicação
- Idle
- FocusSessionActive
- ShortBreakActive
- LongBreakActive
- SessionPausedByUser
- EmergencyBypass
- ErrorDegradedMode

## 2.3 Eventos críticos de domínio
- session.started / session.completed / session.interrupted
- block.applied / block.failed / block.bypassed
- lesson.generated / lesson.delivered / lesson.skipped
- report.updated / report.exported
- permission.missing / permission.granted

---

## 3) Especificação comportamental por feature

## 3.1 Feature A — Timer Pomodoro configurável

### Comportamento esperado
- Usuário define preset com foco, pausa curta, pausa longa e ciclos por bloco.
- Ao iniciar sessão:
  1) valida permissões e política de bloqueio,
  2) aplica bloqueio,
  3) inicia cronômetro de foco.
- Término de foco transita automaticamente para pausa.
- A cada N ciclos completos, pausa longa substitui pausa curta.

### Requisitos funcionais
- RF-A1: Suportar múltiplos presets salvos.
- RF-A2: Exibir estado atual e tempo restante com atualização contínua.
- RF-A3: Persistir estado em caso de reinício do app/sistema.
- RF-A4: Permitir pausa manual e retomada (auditável).

### Requisitos não funcionais
- RNF-A1: Drift de timer máximo aceitável por ciclo (tolerância definida em SLO).
- RNF-A2: Recuperação de estado em abertura do app < 2 segundos.

### Edge cases
- Relógio do sistema alterado durante sessão.
- App fechado inesperadamente.
- Sleep/Wake do macOS no meio do foco.

### Perguntas de validação (senior checklist)
- Vamos usar tempo monotônico interno para evitar quebra por ajuste de relógio?
- Qual regra prevalece ao retomar de sleep: continuar, pausar ou encerrar?

---

## 3.2 Feature B — Bloqueio por apps selecionados

### Comportamento esperado
- Usuário seleciona apps-alvo (bundle identifiers).
- Durante foco, engine monitora processos/janelas e impede uso dos apps bloqueados.
- Ao fim da sessão de foco, bloqueio é removido automaticamente.

### Requisitos funcionais
- RF-B1: Catálogo de apps bloqueáveis com busca e seleção.
- RF-B2: Políticas por contexto (trabalho/estudo).
- RF-B3: Registro de tentativas de abertura e ações tomadas.

### Requisitos de segurança/privacidade
- RF-B4: Solicitar permissões com explicação objetiva de motivo e impacto.
- RF-B5: Não coletar conteúdo do app bloqueado; apenas metadados operacionais.

### Edge cases
- App bloqueado já aberto antes do foco.
- Atualização do app muda bundle id.
- Falha de permissão no meio da sessão.

### Perguntas de validação
- Quando app já estiver aberto, fechar automaticamente ou ocultar/bloquear interação?
- Quais apps nunca podem ser bloqueados por segurança do sistema?

---

## 3.3 Feature C — Bloqueio total temporário

### Comportamento esperado
- Modo crítico de foco com restrição mais rígida de interação.
- Ativação apenas com dupla confirmação + aviso de segurança.
- Deve existir saída de emergência controlada (bypass com justificativa).

### Requisitos funcionais
- RF-C1: Definir janela de bloqueio total atrelada ao timer de foco.
- RF-C2: Fluxo de bypass com cooldown e trilha de auditoria.
- RF-C3: Prevenir lockout irreversível (sempre manter rota de recuperação).

### Requisitos de segurança
- RF-C4: Limites explícitos conforme permissões e políticas do macOS.
- RF-C5: Não interferir em operações críticas do usuário (ex.: chamadas, segurança, acessibilidade assistiva).

### Edge cases
- Usuário precisa encerrar sessão por emergência real.
- Bypass repetido em sequência.

### Perguntas de validação
- Quais critérios diferenciam bypass legítimo vs uso abusivo?
- Qual UX minimiza ansiedade sem enfraquecer foco?

---

## 3.4 Feature D — Pausa com descanso ou microlearning

### Comportamento esperado
- Ao iniciar pausa, usuário pode:
  - descansar sem estímulos, ou
  - consumir microlição no tempo disponível.
- Escolha pode ser automática por preferência do perfil.

### Requisitos funcionais
- RF-D1: Oferecer opção padrão de pausa por perfil.
- RF-D2: Permitir trocar modo no início da pausa.
- RF-D3: Encerrar conteúdo automaticamente quando pausa termina.

### Edge cases
- Lição não carregou a tempo.
- Usuário abre microlearning no final da pausa.

### Perguntas de validação
- Haverá “grace period” de 30-60s para conclusão da lição?
- Em caso de falha de rede, qual conteúdo fallback será exibido?

---

## 3.5 Feature E — Configuração de trilha de microlearning

### Comportamento esperado
- Usuário informa tema, categoria, nível, formato e volume (5/10/15 tópicos ou custom).
- Sistema gera roadmap inicial com módulos e sequência pedagógica.

### Requisitos funcionais
- RF-E1: Wizard de setup em etapas simples (tema → objetivo → profundidade → formato).
- RF-E2: Exibir preview da trilha com possibilidade de ajuste manual.
- RF-E3: Persistir versão da trilha para rastrear mudanças futuras.

### Requisitos de qualidade
- RF-E4: Cada tópico precisa de objetivo de aprendizagem claro.
- RF-E5: Cada tópico deve ter critério mínimo de conclusão.

### Perguntas de validação
- O usuário pode editar ordem dos tópicos ou apenas aceitar sugestão?
- Quantos tópicos iniciais cabem no MVP sem comprometer clareza?

---

## 3.6 Feature F — Pipeline de conteúdo confiável + IA

### Comportamento esperado
Antes de cada pausa com microlearning:
1. sistema identifica próximo objetivo na trilha;
2. busca fontes confiáveis;
3. aplica score de confiança;
4. gera microlição adaptada ao tempo;
5. registra evidências e limitações.

### Requisitos funcionais
- RF-F1: Política de fontes permitidas (whitelist inicial + revisão periódica).
- RF-F2: Score de confiança por item (autoridade, recência, consistência).
- RF-F3: Modo “baixa confiança” com aviso e alternativa.
- RF-F4: Entrega de formato configurado (resumo ou lição com exercício).

### Requisitos de governança
- RF-F5: Transparência de origem do conteúdo para usuário final.
- RF-F6: Proibir conteúdo sem fonte rastreável em modo recomendado.

### Edge cases
- Fontes conflitantes.
- Tema com pouca evidência confiável.
- Resposta do modelo muito longa para tempo de pausa.

### Perguntas de validação
- Qual limiar mínimo de confiança para liberar lição “recomendada”?
- Como versionar prompts/políticas sem quebrar consistência pedagógica?

---

## 3.7 Feature G — Trace/tree da trilha e continuidade

### Comportamento esperado
- Trilha modelada como árvore: Tema → Módulo → Tópico → Microlição.
- Cada nó possui status, tempo acumulado e evidências.
- Seleção da próxima lição considera pré-requisitos e lacunas.

### Requisitos funcionais
- RF-G1: Status por nó (pendente, em andamento, concluído, revisitar).
- RF-G2: Dependências entre tópicos (pré-requisitos).
- RF-G3: Replanejamento incremental sem perder histórico.

### Perguntas de validação
- Como evitar que replanejamento frequente “embaralhe” trilha e gere confusão?
- Quando promover revisão espaçada de um tópico já concluído?

---

## 3.8 Feature H — Relatórios e histórico

### Comportamento esperado
- Dashboard com tempo de foco e estudo por assunto/categoria.
- Histórico diário/semanal/mensal.
- Registro de bypass, interrupções e taxa de conclusão.

### Requisitos funcionais
- RF-H1: Consulta por período e filtros.
- RF-H2: Indicadores de progresso da trilha.
- RF-H3: Exportação futura (post-MVP), mas modelo de dados já preparado.

### Requisitos de privacidade
- RF-H4: Dados locais por padrão; clareza de retenção.
- RF-H5: Opção de exclusão completa de histórico.

### Perguntas de validação
- Quais 5 métricas são realmente acionáveis para usuário comum?
- Como evitar dashboard “bonito, mas irrelevante” para mudança de comportamento?

---

## 4) Fluxos críticos (E2E)

## 4.1 Fluxo 1 — Primeiro uso
1. Onboarding com proposta de valor curta.
2. Pedido de permissões em contexto.
3. Configuração de preset padrão.
4. Configuração opcional da trilha de aprendizado.
5. Primeira sessão guiada.

**Critério de sucesso:** usuário inicia 1º foco em menos de 3 minutos.

## 4.2 Fluxo 2 — Sessão padrão com microlearning
1. Iniciar foco.
2. Aplicar bloqueio.
3. Encerrar foco automaticamente.
4. Gerar e mostrar microlição.
5. Coletar feedback útil/neutro/inútil.
6. Retomar próximo foco.

**Critério de sucesso:** transição foco→pausa sem fricção e sem travamento.

## 4.3 Fluxo 3 — Falha de permissão
1. Sistema detecta ausência/revogação.
2. Entra em modo degradado.
3. Exibe orientação clara para correção.
4. Permite seguir com timer (sem bloqueio) se usuário aceitar.

**Critério de sucesso:** usuário entende impacto e decide conscientemente.

---

## 5) Modelo de dados conceitual (alto nível)

## 5.1 Entidades principais
- UserPreference
- PomodoroPreset
- FocusSession
- BlockPolicy
- BypassEvent
- LearningTrack
- LearningNode
- LessonArtifact
- SourceEvidence
- ConfidenceScore
- DailyReport

## 5.2 Regras de consistência
- Sessão não pode fechar sem status final.
- Lição entregue precisa ter referência de trilha e fonte.
- Evento de bypass sempre vinculado a sessão e motivo.

---

## 6) Requisitos de observabilidade e qualidade

## 6.1 Logs e métricas mínimas
- Latência de transição de estado (foco→pausa).
- Taxa de falha de bloqueio.
- Taxa de lição entregue com confiança alta/média/baixa.
- Uso de CPU/memória por sessão.
- Queda de sessão por erro não tratado.

## 6.2 SLOs iniciais (propostos)
- Disponibilidade do Session Engine: >= 99,5%.
- Sucesso na aplicação de bloqueio por app: >= 98% (com permissões válidas).
- Tempo para exibir microlição em pausa: p95 <= 2s (com cache) / <= 6s (sem cache).

## 6.3 Testabilidade por camada
- Domínio: regras de transição e cálculo de progresso.
- Aplicação: orquestração de fluxos e políticas.
- Infra: permissões, persistência, integração com serviço de conteúdo.
- UX: cenários de estado e mensagens de erro.

---

## 7) Segurança, privacidade e compliance (produto)

## 7.1 Princípios
- Coletar mínimo necessário.
- Explicar claramente permissões sensíveis.
- Dar controle ao usuário sobre dados e retenção.

## 7.2 Controles obrigatórios
- Trilha de auditoria para bypass.
- Segregação entre dados pessoais e telemetria de produto.
- Política de retenção configurável.

## 7.3 Decisões a fechar
- Período padrão de retenção de histórico.
- Estratégia de anonimização para métricas agregadas.

---

## 8) Plano de implementação por fases (com gate de qualidade)

## Fase 1 — Núcleo de foco (4 a 6 semanas)
Entregas:
- Session Engine + presets.
- Bloqueio por app básico.
- Persistência local de sessão.
- Dashboard mínimo de tempo de foco.

Gate de saída:
- Fluxo foco completo estável.
- Recuperação após crash/sleep validada.

## Fase 2 — Microlearning confiável (4 a 6 semanas)
Entregas:
- Setup de trilha.
- Geração de roadmap inicial.
- Entrega de microlições com evidências e score.
- Feedback de utilidade.

Gate de saída:
- Continuidade de trilha sem regressão.
- Conteúdo com rastreabilidade de fonte.

## Fase 3 — Robustez e retenção (3 a 5 semanas)
Entregas:
- Replanejamento incremental por pausa.
- Relatórios avançados de estudo.
- Hardening de permissões e modo degradado.

Gate de saída:
- Estabilidade em cenários críticos.
- KPIs iniciais instrumentados e confiáveis.

---

## 9) Critérios de aceite expandido por risco

## 9.1 Critérios críticos (não negociáveis)
- Nunca iniciar sessão de foco com bloqueio “supostamente ativo” sem validação real.
- Nunca apresentar conteúdo recomendado sem referência rastreável.
- Nunca deixar usuário sem saída segura de bloqueio total.

## 9.2 Critérios de UX
- Usuário entende em até 1 tela: estado atual, tempo restante e próxima ação.
- Mensagens de falha orientam solução, não apenas erro técnico.

## 9.3 Critérios de performance
- Em sessão padrão, uso médio de CPU/memória deve permanecer dentro do orçamento definido em benchmark.

---

## 10) Riscos técnicos priorizados e mitigação
1. **Permissões macOS revogadas em runtime**
   - Mitigação: detector contínuo + fallback orientado.
2. **Bloqueio total com comportamento inconsistente por versão do macOS**
   - Mitigação: matriz de compatibilidade e rollout gradual por versão.
3. **Alucinação/baixa qualidade de conteúdo IA**
   - Mitigação: score de confiança + política de fonte + sinalização explícita.
4. **Complexidade excessiva no onboarding**
   - Mitigação: modo simples default + configurações avançadas progressivas.

---

## 11) Matriz de decisões pendentes (antes de build completo)
- DP-01: Definição final do que significa “bloqueio total” no limite técnico do macOS.
- DP-02: Lista inicial de fontes confiáveis por categoria de assunto.
- DP-03: Política de fallback offline para pausas com microlearning.
- DP-04: Política de bypass (cooldown, limites por dia, impacto em métricas).
- DP-05: Política de retenção e exclusão de dados locais.

---

## 12) Checklist de readiness para iniciar implementação
- [ ] Escopo MVP congelado com critérios de corte.
- [ ] Política de permissões e mensagens UX aprovadas.
- [ ] Modelo de dados revisado por produto + engenharia.
- [ ] Definição de métricas e eventos instrumentáveis.
- [ ] Matriz de risco e plano de fallback aceitos.
- [ ] Plano de testes por camada aprovado.

---

## 13) Conclusão executiva
A implementação deve priorizar confiança operacional: **foco estável, bloqueio previsível e microlearning confiável**.

A principal recomendação arquitetural é iniciar pelo núcleo de sessão e bloqueio com robustez de estado, e só então escalar microlearning com governança de fontes e rastreabilidade. Isso reduz risco técnico, evita retrabalho e preserva a promessa central do produto.
