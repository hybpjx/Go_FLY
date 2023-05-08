// @ts-ignore
(async () => {

    interface Node {
        text: string;
        children: Node[];
    }

    let iRootNode: Node = {} as Node
    let stSeparator: string = "\\"
    const gstPaths: string[] = []


    const loadJson = () => {
        // @ts-ignore
        return new Promise((resolve, reject) => {
            const xhr = new XMLHttpRequest()
            xhr.open('GET', 'dir.json', true)
            xhr.onreadystatechange = () => {
                if (xhr.readyState != 4) {
                    return
                }
                if (xhr.status != 200) {
                    throw 'load Json error'
                }

                iRootNode = JSON.parse(xhr.responseText)
                resolve('')

            }

            xhr.send(null)
        })
    }

    const createDir = (iNode: Node, stParentDir: string) => {
        let path = iNode.text
        if (stParentDir) {
            path = stParentDir + stSeparator + path

        }
        gstPaths.push(path)

    }

    const parseNode = (iNode: Node, stParentDir: string) => {
        if (iNode.text) {
            createDir(iNode, stParentDir)
        }

        if (stParentDir) {
            stParentDir += stSeparator
        }

        iNode.text && (stParentDir += iNode.text)

        if (!iNode.children) {
            return
        }


        for (const iChildNode of iNode.children) {
            parseNode(iChildNode, stParentDir)
        }

    }

    const generateBatFile = () => {
        let stResult = '';
        // for (let path of gstPaths) {
        //     stResult += `md ${path} \n`
        // }

        gstPaths.map(path => {
            stResult += `md ${path} \n`
        })

        const url: string = URL.createObjectURL(new Blob([stResult], {type: 'text/plain'}))

        const domA: HTMLElement = document.createElement('a')

        domA.setAttribute('href', url)
        domA.setAttribute('target', '_bank')
        domA.setAttribute('download', 'generate_dir_03.bat')
        document.body.append(domA)
        domA.click()
        setTimeout(() => {
            document.body.removeChild(domA)
        }, 500)
    }

    await loadJson()

    parseNode(iRootNode, "")
    generateBatFile()
})()