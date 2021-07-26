const Labels = {
    // all
    name: ["Config File Name", "config file name, should be identical"],
    path: ["Executable Path", "path/to/service/executable"],
    args: ["Arguments", "service executable starting arguments string"],
    delay: ["Delay", "delay seconds for starting & stopping service, input like '2,5' "],

    // all services
    svrname: ["Service Name", "name for service"],
    svrid: [
        "Service ID",
        "id for service, leave blank to auto-generate a unique id",
    ],

    // reader
    provider: ["Provider Name", "name of product or system supplying the data"],
    inputfmt: ["Input File Format", "format of input data, one of csv|json"],
    alignmethod: [
        "Align Method",
        "method to align input data to NLPs must be one of prescribed|mapped|inferred",
    ],
    levelmethod: [
        "Level Method",
        "method to apply common scaling this data, one of prescribed|mapped-scale|rules",
    ],
    gencapability: [
        "General Capability",
        "General Capability for assessment results; Literacy or Numeracy",
    ],
    natshost: ["Nats Streaming Host", "hostname/ip of nats broker"],
    natsport: ["Nats Port", "connection port for nats broker"],
    natscluster: ["Nats Cluster Name", "cluster id for nats broker"],
    topic: ["Nats Topic", "nats topic name to publish parsed data items to"],
    folder: ["Watching Folder", "folder to watch for data files"],
    filesuffix: [
        "Only Watch File Suffix",
        "filter files to read by file extension, eg. .csv or .myapp (actual data handling will be determined by input format flag)",
    ],
    interval: ["Interval For Dealing", "watcher poll interval"],
    recursive: ["Dealing With Sub Folder", "watch folders recursively"],
    dotfiles: ["Dealing With Dot Files", "watch dot files"],
    ignore: ["Folders To Be Ignored", "comma separated list of paths to ignore"],
    concurrfiles: [
        "Files' Count In Once Process",
        "pool size for concurrent file processing",
    ],

    // align, level, weight ...
    port: ["Service Port", "current service running port"],

    // align, level ...
    niashost: ["NIAS3 Host", "hostname/ip of nias3"],
    niasport: ["NIAS3 Port", "connection port for nias3"],
    niastoken: ["NIAS3 Token", "token to access nias3"],

    // align
    tchost: ["Text Classifier Host", "text classifier service hostname/ip"],
    tcport: ["Text Classifier Port", "text classifier service connection port"],

    // text classifier

    // weight
    failwhenerr: ["Panic If Error", "if error happens, should service abort?"],

    // hub
    tablename: ["Composite Table Name", "composite table name (.md file)"],

    //////////////////////////////////////////////////////////////

    sel_natsstreaming: ["NatsStreaming", "Select NatsStreaming Config"],
    sel_nias3: ["Nias3", "Select Nias3 Config"],
    sel_benthos_align: ["Benthos for Align", "Select Benthos Config for Align Map Flow"],
    sel_benthos_level: ["Benthos for Level", "Select Benthos Config for Level Map Flow"],
    sel_benthos: ["Benthos for Data", "Select Benthos Config for Data Flow"],
    sel_reader: ["Reader", "Select Reader Config"],
    sel_reader_align_map: ["Reader for Align Map", "Select Reader Config for Align Map"],
    sel_reader_level_map: ["Reader for Level Map", "Select Reader Config for Level Map"],
    sel_align: ["Align", "Select Align Config"],
    sel_txtclassifier: ["TextClassifier", "Select TextClassifier Config"],
    sel_level: ["Level", "Select Level Config"],
    sel_weight: ["Weight", "Select Weight Config"],

    create: ["Create New Config"],
};

export function getLabels() {
    return Labels;
}