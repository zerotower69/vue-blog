import { db } from "@/utils/cloudBase";

export const getData = (dbName: string) =>
  db
    .collection(dbName)
    .get()
    .then(res => res)
    .catch(err => err);
