CREATE TABLE contentMarketing (
    harvestId TEXT NOT NULL,
    commercialPartner TEXT NOT NULL ,
    logoURL TEXT NOT NULL ,
    cerebroScore TEXT NOT NULL,
         UNIQUE (harvestId)
);


CREATE TABLE articles (
    harvestId TEXT NOT NULL ,
    cerebroScore TEXT NOT NULL,
    url TEXT NOT NULL,
    title TEXT NOT NULL,
    cleanImage TEXT NOT NULL,
     UNIQUE (harvestId)
);
