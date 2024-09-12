import { writeFile } from "fs/promises"
const categories = ["business", "person", "owner", "product"]

export const saveFile = async ({category, file, id}) => {
    try {
        const targetDirectory = `/app/storage/${categorie}/${id}_business.${file.extension}`
        //await writeFile(targetDirectory, file)
        return null
    } catch (error) {
        return error
    }
}
