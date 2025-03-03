# zabbix-host-exporter
Automação para Geração de Relatórios de Hosts do Zabbix em CSV

Este projeto tem como objetivo automatizar a coleta de informações sobre todos os hosts cadastrados em uma instância Zabbix e exportar esses dados para um arquivo CSV. Através de uma requisição à API do Zabbix, são extraídos os identificadores e os nomes de cada host registrado. Em seguida, esses dados são organizados e salvos em um arquivo CSV, que pode ser facilmente utilizado para relatórios, auditorias ou análise de infraestrutura.

Funcionalidades:
Conexão com a API Zabbix para obtenção dos dados dos hosts.
Extração dos campos "hostid" e "host" de todos os hosts cadastrados.
Geração automática de um arquivo CSV contendo esses dados.
Simplicidade e praticidade na geração de relatórios de infraestrutura para monitoramento e auditoria.
Esta automação é ideal para equipes de TI e administradores de sistemas que desejam obter relatórios rápidos e bem estruturados sobre os hosts registrados no Zabbix, facilitando a análise e acompanhamento da infraestrutura monitorada.