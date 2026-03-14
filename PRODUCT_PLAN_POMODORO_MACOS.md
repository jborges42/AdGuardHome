# Plano de Produto — App de Pomodoro + Bloqueio Inteligente + Microlearning (macOS)

## 1) Visão do produto
Criar um aplicativo desktop para macOS que combina **foco profundo** (método Pomodoro + bloqueio de distrações) com **aprendizado contínuo** durante os intervalos, via microlearning orientado por IA e fontes confiáveis.

O produto deve ajudar o usuário a:
- Proteger blocos de foco sem fricção.
- Reduzir distrações digitais em momentos críticos.
- Transformar pausas em descanso real **ou** estudo leve, adaptado ao objetivo pessoal.
- Construir trilhas de aprendizado estruturadas com progressão lógica e mensuração de tempo e avanço.

---

## 2) Problema, oportunidade e proposta de valor

### Problema
- Usuários perdem foco por alternância de contexto e acesso fácil a distrações.
- Ferramentas Pomodoro tradicionais não controlam ambiente digital de forma robusta.
- Pausas são subutilizadas quando o usuário quer aprender algo, mas não tem conteúdo curado e progressivo.

### Oportunidade
Unir produtividade + educação em um fluxo único: **foco protegido** e **microlearning com continuidade**, baseado em evidências e rastreabilidade.

### Proposta de valor
> “Um Pomodoro para macOS que protege seu foco de verdade e transforma pausas em aprendizado útil, confiável e personalizado.”

---

## 3) Público-alvo e personas

### Persona 1 — Profissional de conhecimento
- Ex.: PM, dev, designer, analista.
- Dor: interrupções frequentes, baixa consistência de estudo.
- Meta: mais entregas por dia e evolução técnica contínua.

### Persona 2 — Estudante autodidata
- Dor: excesso de conteúdo solto e sem trilha.
- Meta: aprender tópicos de forma progressiva com sessões curtas.

### Persona 3 — Gestor multitarefa
- Dor: agenda fragmentada e cansaço cognitivo.
- Meta: foco com descansos inteligentes e acompanhamento simples.

---

## 4) Objetivos do produto

### Objetivos principais
1. Aumentar aderência a ciclos de foco.
2. Reduzir tempo em apps de distração durante sessões.
3. Melhorar consistência de aprendizado em pausas.
4. Registrar progresso (tempo, assunto, categoria) de forma clara.

### Métricas de sucesso (KPIs)
- Taxa de conclusão de ciclos Pomodoro (%).
- Minutos de foco protegidos por dia.
- Taxa de bloqueios respeitados (sem bypass).
- Minutos de microlearning por semana.
- Conclusão de tópicos da trilha (%).
- Satisfação do conteúdo (feedback rápido: útil/neutro/inútil).

---

## 5) Princípios de UX/UI (macOS-first)
1. **Nativo antes de custom**: seguir Human Interface Guidelines do macOS.
2. **Clareza operacional**: sempre visível o estado atual (foco, pausa, bloqueio, próximo passo).
3. **Baixa carga cognitiva**: poucas decisões por tela, linguagem objetiva.
4. **Acessibilidade**: contraste, teclado, VoiceOver, tamanho de fonte.
5. **Elegância funcional**: sem aparência genérica “de IA”; interface consistente com padrões do sistema.
6. **Progressão orientada**: conteúdo do microlearning mostrado em sequência lógica e com contexto do progresso.

---

## 6) Escopo funcional (MVP + Evoluções)

## 6.1 MVP (obrigatório)
1. **Timer Pomodoro configurável**
   - Duração foco, pausa curta, pausa longa.
   - Número de ciclos até pausa longa.
2. **Bloqueio durante foco**
   - Modo A: bloquear apps selecionados.
   - Modo B: bloqueio total de uso (modo restrito temporário).
3. **Pausa com escolha de modo**
   - Descanso livre.
   - Microlearning.
4. **Microlearning parametrizado**
   - Definição de assunto, categoria e trilha (5, 10, 15+ tópicos).
   - Tipo de material: resumo rápido vs lição mais completa com exercício.
5. **Roadmap inicial por IA**
   - Geração da trilha estruturada antes de iniciar o plano de estudo.
6. **Replanejamento antes de cada pausa**
   - Busca aprofundada e seleção de próximo conteúdo com continuidade.
7. **Rastreio e relatório local**
   - Tempo estudado por assunto/categoria.
   - Histórico de sessões e progresso da trilha.

## 6.2 Pós-MVP
- Sincronização multi-dispositivo.
- Gamificação leve (streaks, badges).
- Recomendação adaptativa por desempenho em exercícios.
- Exportação de relatórios (CSV/PDF).

---

## 7) Regras de negócio (núcleo)

### 7.1 Ciclo Pomodoro
- RN-01: Usuário define tempos padrão e presets personalizados.
- RN-02: Sessão de foco inicia apenas com política de bloqueio ativa.
- RN-03: Encerrado foco, sistema entra em pausa automaticamente.
- RN-04: A cada N ciclos completos, pausa longa obrigatória (configurável).

### 7.2 Bloqueio de distrações
- RN-05: Modo de bloqueio por app impede abertura/uso de apps selecionados durante foco.
- RN-06: Modo bloqueio total restringe interação geral da máquina durante foco (respeitando segurança do sistema).
- RN-07: Bypass só com fluxo explícito (ex.: justificativa + cooldown), registrado em log.
- RN-08: Lista de apps bloqueados é editável fora de sessão ativa.

### 7.3 Microlearning
- RN-09: Usuário pode ativar/desativar microlearning por perfil de estudo.
- RN-10: Ao ativar, usuário define trilha (quantidade de tópicos, profundidade, formato).
- RN-11: IA gera roadmap inicial com ordem pedagógica e objetivos por tópico.
- RN-12: Antes de cada pausa, sistema recalcula próximo bloco considerando progresso + lacunas.
- RN-13: Conteúdo deve indicar nível de confiança da fonte e referência utilizada.
- RN-14: Cada lição deve caber no tempo de pausa configurado.
- RN-15: Se pausa for curta e conteúdo exceder tempo, sistema oferece versão resumida.

### 7.4 Qualidade e confiabilidade do conteúdo
- RN-16: Conteúdo deve priorizar indexadores/fontes confiáveis definidos por política interna.
- RN-17: Resposta gerada pela IA deve incluir “base factual” (resumo + evidências + limites).
- RN-18: Quando houver baixa confiança, sistema sinaliza incerteza e sugere revisão.
- RN-19: Conteúdo não deve inventar referências; sem fonte rastreável, não entra como “recomendado”.

### 7.5 Rastreamento e relatórios
- RN-20: Registrar por sessão: data, duração, modo (foco/pausa), assunto, categoria.
- RN-21: Agregar tempo total por assunto e categoria.
- RN-22: Exibir progresso da trilha (% tópicos concluídos, em andamento, pendentes).
- RN-23: Relatório deve ser salvo localmente e persistir entre reinicializações.

---

## 8) Política de conteúdo com IA (produto)

### Pipeline sugerido (sem detalhar implementação)
1. **Intenção do usuário**: tema, objetivo, nível, formato desejado.
2. **Busca confiável**: coleta em fontes/índices priorizados.
3. **Validação de confiança**: score por recência, autoridade e consistência cruzada.
4. **Síntese didática**: transformação em microlição.
5. **Adaptação temporal**: ajuste para janela da pausa.
6. **Encadeamento pedagógico**: atualização da árvore/trilha de aprendizado.
7. **Feedback loop**: utilidade percebida ajusta próximas lições.

### Estrutura de “trace/tree” da trilha
- Tema principal
  - Módulo
    - Tópico
      - Microlição
      - Exercício (opcional)
      - Evidências/fontes
      - Status (pendente/em andamento/concluído)
      - Tempo acumulado

---

## 9) Histórias de usuário (priorizadas)

### Épico A — Foco e bloqueio
- US-01: Como usuário, quero configurar meus tempos de foco e pausa para adequar ao meu ritmo.
- US-02: Como usuário, quero escolher apps para bloquear durante foco para evitar distrações.
- US-03: Como usuário, quero ativar bloqueio total temporário para sessões críticas.
- US-04: Como usuário, quero saber claramente em qual estado estou (foco/pausa) e quanto tempo falta.

### Épico B — Microlearning
- US-05: Como usuário, quero escolher um tema e montar uma trilha de tópicos para aprender progressivamente.
- US-06: Como usuário, quero definir se prefiro resumos ou lições com exercícios.
- US-07: Como usuário, quero receber uma lição curta durante a pausa, no tempo exato disponível.
- US-08: Como usuário, quero continuidade entre lições para não estudar conteúdos desconexos.

### Épico C — Progresso
- US-09: Como usuário, quero ver quanto tempo estudei por assunto e categoria.
- US-10: Como usuário, quero acompanhar o avanço da minha trilha para manter consistência.
- US-11: Como usuário, quero registrar quando uma lição foi útil para melhorar recomendações.

### Épico D — Experiência macOS
- US-12: Como usuário macOS, quero uma interface nativa, elegante e sem fricção.
- US-13: Como usuário, quero atalhos e interações coerentes com o sistema operacional.

---

## 10) Critérios de aceite (exemplos práticos)

### CA — Timer e fluxo
- CA-01: Dado um preset configurado, quando iniciar sessão, então timer deve seguir tempos definidos.
- CA-02: Ao finalizar foco, pausa inicia automaticamente sem ação manual.

### CA — Bloqueio
- CA-03: Durante foco, apps bloqueados não podem ser utilizados.
- CA-04: No modo total, usuário não acessa recursos não permitidos até término/encerramento autorizado.
- CA-05: Toda quebra de bloqueio deve ser registrada com motivo e horário.

### CA — Microlearning
- CA-06: Com trilha ativa, cada pausa exibe lição alinhada ao próximo tópico pendente.
- CA-07: Se tempo de pausa for 5 min, lição deve caber nesse intervalo.
- CA-08: Conteúdo exibido apresenta fontes/referências e indicador de confiança.

### CA — Relatórios
- CA-09: Após sessões, dashboard mostra tempo total por assunto e categoria.
- CA-10: Progresso da trilha reflete status real de cada tópico.

---

## 11) Jornada do usuário (macrofluxo)
1. Onboarding rápido: objetivo principal (foco, aprendizado ou ambos).
2. Configuração de Pomodoro e bloqueios.
3. (Opcional) Configuração da trilha de microlearning.
4. Sessão de foco com proteção ativa.
5. Pausa: descanso ou microlearning.
6. Retomada automática do próximo ciclo.
7. Encerramento diário com resumo (foco + estudo).

---

## 12) Backlog inicial por fases

### Fase 0 — Descoberta e validação
- Entrevistas com usuários-alvo (10–15).
- Teste de conceito de bloqueio vs aceitação.
- Definição de política de confiabilidade de fontes.

### Fase 1 — MVP operacional
- Timer + estados.
- Bloqueio por app e modo total.
- Configuração inicial de trilha e geração de roadmap.
- Entrega de microlições no intervalo.
- Registro e relatório básico.

### Fase 2 — Qualidade e retenção
- Refino de UX macOS.
- Melhorias de confiabilidade do conteúdo.
- Feedback explícito de utilidade.
- Métricas de engajamento e aprendizado.

### Fase 3 — Escala de valor
- Personalização avançada de trilhas.
- Integrações de calendário/tarefas.
- Relatórios avançados e exportação.

---

## 13) Riscos e mitigação
- **Risco técnico (bloqueio no macOS)**: limitações de permissões e segurança.
  - Mitigação: escopo claro de capacidades, onboarding de permissões transparente.
- **Risco de qualidade de conteúdo (IA)**: alucinação ou baixa confiabilidade.
  - Mitigação: política de fontes + score de confiança + sinalização de incerteza.
- **Risco de rejeição por rigidez**: bloqueio muito agressivo.
  - Mitigação: perfis de rigidez e bypass controlado.
- **Risco de UX complexa**: excesso de configuração inicial.
  - Mitigação: presets inteligentes + modo avançado opcional.

---

## 14) Requisitos não-funcionais (produto)
- Performance percebida: ações principais em até poucos cliques e resposta imediata.
- Confiabilidade: sessões não podem “perder estado”.
- Privacidade: dados de estudo e uso com transparência e controle do usuário.
- Observabilidade: eventos de sessão, bloqueio e aprendizado rastreáveis.
- Acessibilidade: compatível com boas práticas do ecossistema macOS.

---

## 15) Definição de pronto (DoD) para histórias de produto
Uma história é considerada pronta quando:
1. Regra de negócio validada com PO/UX.
2. Critérios de aceite testáveis e sem ambiguidade.
3. Estados de erro e edge cases definidos.
4. Impacto em métricas/KPIs identificado.
5. Experiência alinhada a padrões nativos macOS.
6. Telemetria necessária especificada.

---

## 16) Próximos passos práticos (sem código)
1. Priorizar backlog do MVP com matriz impacto x esforço.
2. Criar mapa de telas e fluxos (wireframes de baixa fidelidade).
3. Definir política editorial de conteúdo + critérios de confiabilidade.
4. Escrever PRD v1 com escopo fechado de 8–12 semanas.
5. Planejar ciclo de validação com usuários (antes e após MVP).

---

## 17) Resumo executivo
Este produto tem potencial forte por atacar dois problemas reais em conjunto: **foco** e **aprendizado contínuo**. A chave do sucesso será equilibrar:
- bloqueio eficaz sem frustração,
- conteúdo útil com confiança rastreável,
- experiência nativa e elegante no macOS.

Se esse equilíbrio for mantido, o app deixa de ser “apenas um timer” e vira um **sistema pessoal de foco e evolução profissional**.
