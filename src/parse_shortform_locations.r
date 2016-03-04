#!/usr/bin/env r

# littler script (piped to by `cut` in parse_shortform_locations.sh)

library('magrittr')

position.field.names <- list('SuccessfullyMapped',
                          'MappedSeqInfo.GenBankID',
                          'MappedSeqInfo.RefSeqID',
                          'MappedSeqInfo.SequenceType',
                          'ChrStart',
                          'ChrEnd',
                          'MappedVersionNumbers',
                          'MappedVersionAccessions',
                          'Accession1Method',
                          'Accession2Method',
                          'AssemblyName',
                          'GenBankAssemblyAcc',
                          'RefSeqAssemblyAcc',
                          'AssemblyStatus')

ExpandLocationPositions <- function(location) {
  
  # a new row in the data frame (and thus table output) for each item in this list (corresponding to <position> tags per issue <location>)
  position.fields <- lapply(strsplit(location, "::@@::"),
                            function(x) { strsplit(x, ":@:") })[[1]]
  lapply(position.fields, function(x) {
    c(x[1:6],
      lapply(strsplit(x[7], ':::')[[1]] %>% strsplit(":@@@:"), function(x){ x[1] }) %>% unlist %>% paste0(collapse=","),
      lapply(strsplit(x[7], ':::')[[1]] %>% strsplit(":@@@:"), function(x){ x[2] }) %>% unlist %>% paste0(collapse=","),
      x[8:13])
  }) %>%
    lapply(function(x) {
      positions.list <- vector("list", length(x)) %>%
        setNames(position.field.names)
      for (i in 1:length(x)) { positions.list[position.field.names[[i]]] = x[i] }
      return(positions.list %>% as.data.frame)
    }) %>%
    do.call("rbind", .)
}

ExpandLocation <- function(df.row) {
  new.table.names <- c('id','chr', position.field.names %>% unlist)
  new.row <- vector("list", length(new.table.names)) %>%
    setNames(new.table.names)
  expanded <- df.row[['location']] %>% ExpandLocationPositions
  lapply(position.field.names, function(x) { new.row[[x]] <<- expanded[[x]] })
  lapply(list('id','chr'), function(x) { new.row[[x]] <<- rep(df.row[[x]], nrow(expanded)) })
  return(new.row %>% as.data.frame)
}

grc.table <- read.table(file = file("stdin"), sep = '\t',
			stringsAsFactors = FALSE, header = TRUE)
colnames(grc.table) <- c('id', 'chr', 'location')

output.table <- apply(grc.table, 1, ExpandLocation) %>%
  do.call("rbind", .)

write.table(output.table, file = "", sep = '\t', quote = FALSE, row.names = FALSE)