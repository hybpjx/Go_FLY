"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
// @ts-ignore
(() => __awaiter(void 0, void 0, void 0, function* () {
    let iRootNode = {};
    let stSeparator = "\\";
    const gstPaths = [];
    const loadJson = () => {
        // @ts-ignore
        return new Promise((resolve, reject) => {
            const xhr = new XMLHttpRequest();
            xhr.open('GET', 'dir.json', true);
            xhr.onreadystatechange = () => {
                if (xhr.readyState != 4) {
                    return;
                }
                if (xhr.status != 200) {
                    throw 'load Json error';
                }
                iRootNode = JSON.parse(xhr.responseText);
                resolve('');
            };
            xhr.send(null);
        });
    };
    const createDir = (iNode, stParentDir) => {
        let path = iNode.text;
        if (stParentDir) {
            path = stParentDir + stSeparator + path;
        }
        gstPaths.push(path);
    };
    const parseNode = (iNode, stParentDir) => {
        if (iNode.text) {
            createDir(iNode, stParentDir);
        }
        if (stParentDir) {
            stParentDir += stSeparator;
        }
        iNode.text && (stParentDir += iNode.text);
        if (!iNode.children) {
            return;
        }
        for (const iChildNode of iNode.children) {
            parseNode(iChildNode, stParentDir);
        }
    };
    const generateBatFile = () => {
        let stResult = '';
        // for (let path of gstPaths) {
        //     stResult += `md ${path} \n`
        // }
        gstPaths.map(path => {
            stResult += `md ${path} \n`;
        });
        const url = URL.createObjectURL(new Blob([stResult], { type: 'text/plain' }));
        const domA = document.createElement('a');
        domA.setAttribute('href', url);
        domA.setAttribute('target', '_bank');
        domA.setAttribute('download', 'generate_dir_03.bat');
        document.body.append(domA);
        domA.click();
        setTimeout(() => {
            document.body.removeChild(domA);
        }, 500);
    };
    yield loadJson();
    parseNode(iRootNode, "");
    generateBatFile();
}))();
