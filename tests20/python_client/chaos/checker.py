import datetime
from enum import Enum
from random import randint

from time import sleep
from base.collection_wrapper import ApiCollectionWrapper
from common import common_func as cf
from common import common_type as ct
import constants
from utils.util_log import test_log as log


class Op(Enum):
    create = 'create'
    insert = 'insert'
    flush = 'flush'
    index = 'index'
    search = 'search'
    query = 'query'

    unknown = 'unknown'


class Checker:
    def __init__(self):
        self._succ = 0
        self._fail = 0
        self._running = True
        self.c_wrap = ApiCollectionWrapper()
        self.c_wrap.init_collection(name=cf.gen_unique_str('Checker_'),
                                    schema=cf.gen_default_collection_schema())
        self.c_wrap.insert(data=cf.gen_default_list_data(nb=constants.ENTITIES_FOR_SEARCH),
                           check_task='check_nothing')
        self.initial_entities = self.c_wrap.num_entities    # do as a flush

    def total(self):
        return self._succ + self._fail

    def succ_rate(self):
        return self._succ / self.total() if self.total() != 0 else 0

    def terminate(self):
        self._running = False

    def reset(self):
        self._succ = 0
        self._fail = 0


class SearchChecker(Checker):
    def __init__(self):
        super().__init__()
        self.c_wrap.load()   # do load before search

    def keep_running(self):
        while self._running is True:
            search_vec = cf.gen_vectors(5, ct.default_dim)
            _, result = self.c_wrap.search(
                                data=search_vec,
                                anns_field=ct.default_float_vec_field_name,
                                param={"nprobe": 32},
                                limit=1, check_task='check_nothing'
                            )
            if result:
                self._succ += 1
            else:
                self._fail += 1
            sleep(constants.WAIT_PER_OP / 10)


class InsertFlushChecker(Checker):
    def __init__(self, flush=False):
        super().__init__()
        self._flush = flush

    def keep_running(self):
        while self._running:
            init_entities = self.c_wrap.num_entities
            _, insert_result = \
                self.c_wrap.insert(data=cf.gen_default_list_data(nb=constants.DELTA_PER_INS),
                                   check_task='check_nothing')
            if not self._flush:
                if insert_result:
                    self._succ += 1
                else:
                    self._fail += 1
                sleep(constants.WAIT_PER_OP / 10)
            else:
                if self.c_wrap.num_entities == (init_entities + constants.DELTA_PER_INS):
                    self._succ += 1
                else:
                    self._fail += 1


class CreateChecker(Checker):
    def __init__(self):
        super().__init__()

    def keep_running(self):
        while self._running is True:
            _, result = self.c_wrap.init_collection(
                                    name=cf.gen_unique_str("CreateChecker_"),
                                    schema=cf.gen_default_collection_schema(),
                                    check_task='check_nothing'
                                )
            if result:
                self._succ += 1
                self.c_wrap.drop(check_task="check_nothing")
            else:
                self._fail += 1
            sleep(constants.WAIT_PER_OP / 10)


class IndexChecker(Checker):
    def __init__(self):
        super().__init__()
        self.c_wrap.insert(data=cf.gen_default_list_data(nb=5*constants.ENTITIES_FOR_SEARCH),
                           check_task='check_nothing')
        log.debug(f"Index ready entities: {self.c_wrap.num_entities }")  # do as a flush before indexing

    def keep_running(self):
        while self._running:
            _, result = self.c_wrap.create_index(ct.default_float_vec_field_name,
                                                 constants.DEFAULT_INDEX_PARAM,
                                                 name=cf.gen_unique_str('index_'),
                                                 check_task='check_nothing')
            if result:
                self._succ += 1
                self.c_wrap.drop_index(check_task='check_nothing')
            else:
                self._fail += 1


class QueryChecker(Checker):
    def __init__(self):
        super().__init__()
        self.c_wrap.load()      # load before query

    def keep_running(self):
        while self._running:
            int_values = []
            for _ in range(5):
                int_values.append(randint(0, constants.ENTITIES_FOR_SEARCH))
            term_expr = f'{ct.default_int64_field_name} in {int_values}'
            _, result = self.c_wrap.query(term_expr, check_task='check_nothing')
            if result:
                self._succ += 1
            else:
                self._fail += 1
            sleep(constants.WAIT_PER_OP / 10)

#
# if __name__ == '__main__':
#     from pymilvus_orm import connections
#     connections.add_connection(default={"host": '10.98.0.7', "port": 19530})
#     conn = connections.connect(alias='default')
#     c_w = ApiCollectionWrapper()
#     c_w.init_collection(name=cf.gen_unique_str("collection_4_search_"),
#                         schema=cf.gen_default_collection_schema())
#     c_w.insert(data=cf.gen_default_list_data(nb=constants.ENTITIES_FOR_SEARCH))
#     log.debug(f"nums: {c_w.num_entities}")
#     # c_w.load()
#     # # int_values = []
#     # # for _ in range(5):
#     # #     int_values.append(randint(0, constants.ENTITIES_FOR_SEARCH))
#     # term_expr = f'{ct.default_int64_field_name} in [1,2,3,4,5]'
#     # log.debug(term_expr)
#     # res, result = c_w.query(term_expr)
#
#     res, result = c_w.create_index(ct.default_float_vec_field_name,
#                                    constants.DEFAULT_INDEX_PARAM,
#                                    name=cf.gen_unique_str('index_'),
#                                    check_task='check_nothing')
#     log.debug(res)
